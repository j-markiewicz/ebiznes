plugins {
	alias(libs.plugins.kotlin.jvm)
	alias(libs.plugins.kotlinx.serialization)
	alias(libs.plugins.ktor)
	application
}

group = "example.ebiznes"
version = "0.0.1"

application {
	mainClass.set("example.ebiznes.MainKt")
}

kotlin {
	jvmToolchain(21)
}

dependencies {
	implementation(libs.kotlinx.serialization.json)
	implementation(libs.logback.classic)
	implementation(libs.ktor.client.core)
	implementation(libs.ktor.client.cio)
	implementation(libs.ktor.client.websockets)
	implementation(libs.dotenv)
}
