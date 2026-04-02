import example.ebiznes.ebiz2._
import org.scalatra._
import jakarta.servlet.ServletContext

class ScalatraBootstrap extends LifeCycle {
	override def init(context: ServletContext) = {
		context.mount(new Products, "/products")
		context.mount(new Cart, "/cart")
		context.mount(new Categories, "/categories")
		context.setInitParameter("org.scalatra.cors.allowedOrigins", "http://localhost:8080,https://ebiznes.example")
		context.setInitParameter("org.scalatra.cors.allowedMethods", "GET,POST,PUT,PATCH,DELETE")
		context.setInitParameter("org.scalatra.cors.allowCredentials", "false")
	}
}
