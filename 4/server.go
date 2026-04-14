package main

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())

	e.GET("/", index)

	e.GET("/products", listProducts)
	e.POST("/products", createProduct)
	e.GET("/products/:id", readProduct)
	e.PATCH("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

// GET /
func index(c *echo.Context) error {
	return c.HTML(http.StatusOK, `
		<form action="/products" method="POST">
			<label>Nazwa: <input type="text" name="name" required /></label>
			<label>Opis: <input type="text" name="description" required /></label>
			<label>Cena: <input type="number" name="price" required /></label>
			<input type="submit" value="Wyślij" />
		</form>
	`)
}

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       uint32 `json:"price"`
}

type Products struct {
	products map[uuid.UUID]Product
}

func newProducts() Products {
	return Products{make(map[uuid.UUID]Product)}
}

func (p *Products) exists(id uuid.UUID) bool {
	_, exists := p.products[id]
	return exists
}

func (p *Products) list() map[uuid.UUID]Product {
	return p.products
}

func (p *Products) create(product Product) (uuid.UUID, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return id, err
	}

	p.products[id] = product
	return id, nil
}

func (p *Products) read(id uuid.UUID) Product {
	return p.products[id]
}

func (p *Products) update(id uuid.UUID, product Product) {
	p.products[id] = product
}

func (p *Products) delete(id uuid.UUID) {
	delete(p.products, id)
}

var products = newProducts()

// GET /products
func listProducts(c *echo.Context) error {
	return c.JSON(http.StatusOK, products.list())
}

// POST /products
func createProduct(c *echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	priceStr := c.FormValue("price")

	if name == "" || description == "" || priceStr == "" {
		return c.String(http.StatusBadRequest, "required value missing")
	}

	price, err := strconv.ParseUint(priceStr, 10, 32)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	product := Product{name, description, uint32(price)}
	id, err := products.create(product)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Location", "/products/"+id.String())
	return c.JSON(http.StatusCreated, product)
}

// GET /products/:id
func readProduct(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	if !products.exists(id) {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, products.read(id))
}

// PATCH /products/:id
func updateProduct(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	if !products.exists(id) {
		return c.NoContent(http.StatusNotFound)
	}
	product := products.read(id)

	name := c.FormValue("name")
	if name != "" {
		product.Name = name
	}

	description := c.FormValue("description")
	if description != "" {
		product.Description = description
	}

	priceStr := c.FormValue("price")
	if priceStr != "" {
		price, err := strconv.ParseUint(priceStr, 10, 32)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		product.Price = uint32(price)
	}

	products.update(id, product)
	return c.JSON(http.StatusOK, product)
}

// DELETE /products/:id
func deleteProduct(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	if !products.exists(id) {
		return c.NoContent(http.StatusNotFound)
	}

	products.delete(id)
	return c.NoContent(http.StatusOK)
}
