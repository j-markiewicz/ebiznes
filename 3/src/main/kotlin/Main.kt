package example.ebiznes

import kotlinx.serialization.json.*
import kotlinx.serialization.*
import kotlinx.coroutines.*
import kotlinx.coroutines.channels.*
import io.ktor.http.*
import io.ktor.websocket.*
import io.ktor.client.*
import io.ktor.client.engine.cio.*
import io.ktor.client.plugins.websocket.*
import io.ktor.client.request.*
import io.ktor.client.statement.*
import io.github.cdimascio.dotenv.dotenv
import java.util.concurrent.Executors

@Serializable
data class Event(val op: Int, val t: String? = null, val s: Int? = null, val d: JsonElement? = null)

data class Product(val name: String, val description: String, val price: Int)

class Sequence(var value: Int? = null) {
	fun get(): Int? {
		return value
	}
	
	fun set(new_value: Int?) {
		if (new_value != null) {
			value = new_value
		}
	}
}

val products = mapOf(
	"owoce" to listOf(
		Product("Truskawki", "opakowanie 500g", 699)
	),
	"warzywa" to listOf(
		Product("Ogórki", "1kg, świeże", 899),
		Product("Marchewki", "1kg luzem", 399)
	),
)

suspend fun main() {
	val dotenv = dotenv()
	val api_base = dotenv["API_BASE"] ?: "https://discord.com/api/v10"
	val client_secret = dotenv["CLIENT_SECRET"]
	val bot_url = dotenv["BOT_URL"] ?: "https://github.com/j-markiewicz/ebiznes/tree/main/3"
	val client = HttpClient(CIO) {
		install(WebSockets) {
			pingIntervalMillis = 10_000
		}
	}

	println("Bot starting with api base $api_base as Bot ${client_secret.subSequence(0, client_secret.indexOfLast {it == '.'} + 4)}...")

	val headers = listOf(
		listOf("Accept", "application/json"),
		listOf("User-Agent", "MessageBot ($bot_url, v0.0.1)"),
		listOf("Authorization", "Bot $client_secret")
	)

	val bot_id = test_request(client, headers, api_base)
	val gateway_url = get_gateway_url(client, headers, api_base)
	val (tx, rx, seq) = connect_gateway(client, headers, gateway_url, client_secret)

	while (true) {
		val frame = rx.receive()
		assert(frame.fin)
		assert(frame is Frame.Text)
		val event = Json.decodeFromString<Event>((frame as Frame.Text).readText())
		seq.set(event.s)

		when (event.op) {
			0 -> when (event.t) {
				"GUILD_CREATE" -> {
					val name = event.d?.jsonObject?.get("name")?.jsonPrimitive?.content
					println("Bot is in server \"$name\"")
				}
				"MESSAGE_CREATE" -> {
					val content = event.d?.jsonObject?.get("content")?.jsonPrimitive?.content
					val username = event.d?.jsonObject?.get("author")?.jsonObject?.get("username")?.jsonPrimitive?.content
					val is_bot = event.d?.jsonObject?.get("author")?.jsonObject?.get("bot")?.jsonPrimitive?.booleanOrNull ?: false
					val channel_id = event.d?.jsonObject?.get("channel_id")?.jsonPrimitive?.content
					val mentions = event.d?.jsonObject?.get("mentions")?.jsonArray?.map {
						it.jsonObject?.get("id")?.jsonPrimitive?.content
					} ?: listOf()

					if (is_bot) {
						println("Bot received message \"$content\" from bot \"$username\"")
					} else if (channel_id == null) {
						println("Bot received message \"$content\" from \"$username\" on unknown channel")
					} else if (content != null && mentions.contains(bot_id)) {
						println("Bot received message \"$content\" from \"$username\"")

						var sent = false

						if (content.contains("list")) {
							val categories = products.keys.map {"*$it*"}.joinToString()
							send_message(client, headers, api_base, channel_id, "Product categories: $categories")
							sent = true
						}

						for (category in products.keys) {
							if (content.contains(category)) {
								val products = products[category]?.map {"*${it.name}* (${it.price}¤) - ${it.description}"}?.joinToString("\n") ?: "*no products in this category*"
								send_message(client, headers, api_base, channel_id, "Products in *$category*:\n$products")
								sent = true
							}
						}

						if (!sent) {
							send_message(client, headers, api_base, channel_id, "Write *list* to get a list of product categories\nWrite the name of a category to get a list of its products")
						}
					} else {
						println("Bot received message \"$content\" from \"$username\", in which it wasn't tagged")
					}
				}
				else -> println("Received dispatch event of type ${event.t}: ${event.d}")
			}
			1 -> send_heartbeat(tx, seq.get())
			7 -> {
				println("Bot got disconnected")
				return
			}
			9 -> {
				println("Bot got deauthenticated")
				return
			}
			11 -> println("Received heartbeat ACK")
			else -> println("Received other event: ${event}")
		}
	}
}

suspend fun test_request(client: HttpClient, headers: List<List<String>>, api_base: String): String {
	val res = client.get("$api_base/oauth2/applications/@me") {
		headers {
			for (header in headers) {
				append(header[0], header[1])
			}
		}
	}

	if (res.status.isSuccess()) {
		val json = Json.decodeFromString<JsonObject>(res.bodyAsText())
		val id = json["bot"]?.jsonObject?.get("id")?.jsonPrimitive?.content
		val name = json["bot"]?.jsonObject?.get("username")?.jsonPrimitive?.content
		val discr = json["bot"]?.jsonObject?.get("discriminator")?.jsonPrimitive?.content
		println("Test request successful, authenticated as $name#$discr ($id)")

		if (id == null || name == null || discr == null) {
			throw RuntimeException("API test request returned invalid data")
		}

		return id
	} else {
		println("Test request unsuccessful: ${res.status}")
		throw RuntimeException("API test request unsuccessful: ${res.status}")
	}
}

suspend fun get_gateway_url(client: HttpClient, headers: List<List<String>>, api_base: String): String {
	val res = client.get("$api_base/gateway/bot") {
		headers {
			for (header in headers) {
				append(header[0], header[1])
			}
		}
	}

	if (res.status.isSuccess()) {
		val json = Json.decodeFromString<JsonObject>(res.bodyAsText())
		val url = json["url"]?.jsonPrimitive?.content

		if (url == null) {
			throw RuntimeException("API gateway url request returned invalid data: ${res.bodyAsText()}")
		}

		return url
	} else {
		throw RuntimeException("API gateway url request unsuccessful: ${res.status}")
	}
}

suspend fun connect_gateway(client: HttpClient, headers: List<List<String>>, gateway_url: String, client_secret: String): Triple<SendChannel<Frame>, ReceiveChannel<Frame>, Sequence> {
	val session = client.webSocketSession("$gateway_url?v=10&encoding=json") {
		headers {
			for (header in headers) {
				append(header[0], header[1])
			}
		}
	}

	val tx = session.outgoing
	val rx = session.incoming
	val last_seq = Sequence()

	println("Connected to gateway at $gateway_url, awaiting server hello")

	var frame = rx.receive()
	assert(frame.fin)
	assert(frame is Frame.Text)
	var event = Json.decodeFromString<Event>((frame as Frame.Text).readText())
	assert(event.op == 10)
	val heartbeat_interval = event.d?.jsonObject?.get("heartbeat_interval")?.jsonPrimitive?.intOrNull ?: 30000

	Executors.newSingleThreadExecutor().execute { runBlocking {
		delay((heartbeat_interval * (500..1000).random() / 1000).toLong())

		while (true) {
			send_heartbeat(tx, last_seq.get())
			delay((heartbeat_interval).toLong())
		}
	}}

	println("Received server hello, authenticating")

	tx.send(Frame.Text(Json.encodeToString(Event(op = 2, d = JsonObject(mapOf(
		"token" to JsonPrimitive(client_secret),
		"intents" to JsonPrimitive(37377),
		"properties" to JsonObject(mapOf(
			"os" to JsonPrimitive("linux"),
			"browser" to JsonPrimitive("MessageBot"),
			"device" to JsonPrimitive("MessageBot server")
		)),
	))))))

	println("Sent identify event, awaiting server ready")

	frame = rx.receive()
	assert(frame.fin)
	assert(frame is Frame.Text)
	event = Json.decodeFromString<Event>((frame as Frame.Text).readText())
	assert(event.op == 0)
	assert(event.t == "READY")
	last_seq.set(event.s)

	println("Connected and authenticated to gateway at $gateway_url, bot ready")

	return Triple(tx, rx, last_seq)
}

suspend fun send_heartbeat(tx: SendChannel<Frame>, last_seq: Int?) {
	println("Sending heartbeat")
	tx.send(Frame.Text(Json.encodeToString(Event(op = 1, d = JsonPrimitive(last_seq)))))
}

suspend fun send_message(client: HttpClient, headers: List<List<String>>, api_base: String, channel_id: String, content: String) {
	val res = client.post("$api_base/channels/$channel_id/messages") {
		setBody(Json.encodeToString(mapOf(
			"content" to JsonPrimitive(content),
			"flags" to JsonPrimitive(4),
		)))
		headers {
			append("Content-Type", "application/json")
			for (header in headers) {
				append(header[0], header[1])
			}
		}
	}

	if (!res.status.isSuccess()) {
		println("Message sent unsuccessfully: ${res.status}")
		throw RuntimeException("API call to send message unsuccessful: ${res.status}")
	}
}
