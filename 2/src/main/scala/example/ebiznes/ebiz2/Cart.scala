package example.ebiznes.ebiz2

import scala.collection.mutable;
import org.scalatra._
import org.json4s._
import org.json4s.native.JsonMethods._
import org.json4s.native.Serialization
import org.json4s.native.Serialization.{read, write}
import org.json4s.JsonDSL.WithBigDecimal._

class Item(var id: Int, var amount: Int)

object CartDb {
	var all: mutable.HashMap[Int, Int] = mutable.HashMap()

	def list() = {
		for (id, amount) <- CartDb.all
			yield Item(id, amount)
	}

	def add(item: Item) = {
		if all.contains(item.id) then
			all(item.id) += item.amount
		else
			all.addOne(item.id -> item.amount)

		if all(item.id) < 0 || all(item.id) == 0 then
			all.remove(item.id)
	}

	def remove(item: Item) = add(Item(item.id, -item.amount))
}

class Cart extends ScalatraFilter with CorsSupport {
	before() {
		contentType = "application/json"
	}

	get("/cart") {
		write(CartDb.list())
	}
	
	post("/cart/add") {
		try
			CartDb.add(parse(request.body).extract[Item])
		catch
			case _ => halt(400, s"invalid request body")
	}
	
	delete("/cart/remove") {
		try
			CartDb.remove(parse(request.body).extract[Item])
		catch
			case _ => halt(400, s"invalid request body")
	}
}
