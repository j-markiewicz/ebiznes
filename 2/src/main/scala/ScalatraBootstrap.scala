import example.ebiznes.ebiz2._
import org.scalatra._
import jakarta.servlet.ServletContext

class ScalatraBootstrap extends LifeCycle {
	override def init(context: ServletContext) = {
		context.mount(new Ebiznes, "/")
	}
}
