package main

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

var db Database

func main() {
	var err error
	db, err = NewDatabase("file::memory:")
	if err != nil {
		panic("failed to connect database:\n" + err.Error())
	}

	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", Index)

	e.GET("/products", ListProducts)
	e.POST("/products", CreateProduct)
	e.GET("/products/:id", ReadProduct)
	e.PATCH("/products/:id", UpdateProduct)
	e.DELETE("/products/:id", DeleteProduct)

	if err := e.Start(":8000"); err != nil {
		panic("failed to start server:\n" + err.Error())
	}
}
