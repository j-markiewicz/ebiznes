package example.ebiznes.ebiz2

import org.scalatra._
import org.json4s._

class Product(var name: String, var description: String, var price: Int)

object Products {
	var all = Map(
		1 -> Product("Ogórki", "1kg, świeże", 899),
		2 -> Product("Marchewki", "1kg luzem", 399),
		3 -> Product("Truskawki", "opakowanie 500g", 699),
	)
}

class Ebiznes extends ScalatraFilter {
	get("/products") {
		contentType = "application/json"

		s"[${
			Products.all.values.map(
				p => s"{\"name\": \"${p.name}\", \"description\": \"${p.description}\", \"price\": ${p.price}}"
			).mkString(", ")
		}]"
	}
}
