package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/products", listProducts)
	e.GET("/products/:id", readProduct)

	e.POST("/pay", pay)

	if err := e.Start(":8000"); err != nil {
		panic("failed to start server:\n" + err.Error())
	}
}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
}

var products = map[string]Product{
	"1": {"Truskawki", "opakowanie 500g", 699},
	"2": {"Ogórki", "1kg, świeże", 899},
	"3": {"Marchewki", "1kg luzem", 399},
}

// GET /products
func listProducts(c *echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

// GET /products/:id
func readProduct(c *echo.Context) error {
	return c.JSON(http.StatusOK, products[c.Param("id")])
}

// POST /pay
func pay(c *echo.Context) error {
	time.Sleep(1_500_000_000)
	return c.NoContent(http.StatusOK)
}
