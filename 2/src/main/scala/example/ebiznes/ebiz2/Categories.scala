package example.ebiznes.ebiz2

import scala.collection.mutable;
import org.scalatra._
import org.json4s._
import org.json4s.native.JsonMethods._
import org.json4s.native.Serialization
import org.json4s.native.Serialization.{read, write}
import org.json4s.JsonDSL.WithBigDecimal._

class CategoryAction(var action: String, var product: Int)

object CategoriesDb {
	var all = mutable.HashMap(
		"Warzywa" -> mutable.HashSet(1, 2),
		"Owoce" -> mutable.HashSet(3),
	)

	def list() = {
		CategoriesDb.all.keys
	}

	def products(category: String) = {
		if all.contains(category) then
			all(category)
		else
			mutable.HashSet()
	}

	def add(category: String, product: Int) = {
		if all.contains(category) then
			all(category).add(product)
		else
			all.addOne(category -> mutable.HashSet(product))
	}

	def remove(category: String, product: Int) = {
		if all.contains(category) then
			all(category).remove(product)
	}
}

class Categories extends ScalatraFilter with CorsSupport {
	before() {
		contentType = "application/json"
	}

	get("/categories") {
		write(CategoriesDb.list())
	}

	get("/categories/:name") {
		write(CategoriesDb.products(params("name")))
	}
	
	patch("/categories/:name") {
		val ca = parse(request.body).extract[CategoryAction]
		ca.action match
			case "add" => CategoriesDb.add(params("name"), ca.product)
			case "remove" => CategoriesDb.remove(params("name"), ca.product)
			case _ => halt(400, s"invalid action")
	}
}
