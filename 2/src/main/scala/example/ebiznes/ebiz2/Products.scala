package example.ebiznes.ebiz2

import scala.collection.mutable;
import org.scalatra._
import org.json4s._
import org.json4s.native.JsonMethods._
import org.json4s.native.Serialization
import org.json4s.native.Serialization.{read, write}
import org.json4s.JsonDSL.WithBigDecimal._

implicit val formats: Formats = Serialization.formats(NoTypeHints)

class Product(var name: String, var description: String, var price: Int)
class IdProduct(var id: Int, var name: String, var description: String, var price: Int)

object ProductsDb {
	var all = mutable.HashMap(
		1 -> Product("Ogórki", "1kg, świeże", 899),
		2 -> Product("Marchewki", "1kg luzem", 399),
		3 -> Product("Truskawki", "opakowanie 500g", 699),
	)

	def idProducts() = {
		for (id, product) <- ProductsDb.all
			yield IdProduct(id, product.name, product.description, product.price)
	}

	def idProduct(id: Int) = {
		val product = all(id)
		IdProduct(id, product.name, product.description, product.price)
	}

	def remove(id: Int) = {
		all.remove(id)
	}

	def add(product: Product) = {
		all += (all.keys.max + 1 -> product)
	}

	def set(id: Int, product: Product) = {
		all(id) = product
	}
}

class Products extends ScalatraFilter with CorsSupport {
	before() {
		contentType = "application/json"
	}

	get("/products") {
		write(ProductsDb.idProducts())
	}

	get("/products/:id") {
		try
			write(ProductsDb.idProduct(params("id").toInt))
		catch
			case _ => halt(404, s"product ${params("id")} not found")
	}
	
	post("/products") {
		try
			ProductsDb.add(parse(request.body).extract[Product])
		catch
			case _ => halt(400, s"invalid request body")
	}
	
	put("/products/:id") {
		try
			ProductsDb.set(params("id").toInt, parse(request.body).extract[Product])
		catch
			case e: NumberFormatException => halt(404, s"product ${params("id")} not found")
			case e: NoSuchElementException => halt(404, s"product ${params("id")} not found")
			case _ => halt(400, s"invalid request body")
	}
	
	delete("/products/:id") {
		try
			ProductsDb.remove(params("id").toInt)
		catch
			case _ => halt(404, s"product ${params("id")} not found")
	}
}
