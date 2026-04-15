package main

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

// GET /
func Index(c *echo.Context) error {
	return c.HTML(http.StatusOK, `
		<form action="/products" method="POST">
			<label>Nazwa: <input type="text" name="name" required /></label>
			<label>Opis: <input type="text" name="description" required /></label>
			<label>Cena: <input type="number" name="price" required /></label>
			<input type="submit" value="Wyślij" />
		</form>
	`)
}

// GET /products
func ListProducts(c *echo.Context) error {
	res, err := db.List()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// POST /products
func CreateProduct(c *echo.Context) error {
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

	product := Product{Name: name, Description: description, Price: uint32(price)}
	id, err := db.Create(product)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Location", "/products/"+id.String())
	return c.JSON(http.StatusCreated, product)
}

// GET /products/:id
func ReadProduct(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	exists, err := db.Exists(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.NoContent(http.StatusNotFound)
	}

	res, err := db.Read(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// PATCH /products/:id
func UpdateProduct(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	exists, err := db.Exists(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.NoContent(http.StatusNotFound)
	}

	product, err := db.Read(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

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

	err = db.Update(id, product)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, product)
}

// DELETE /products/:id
func DeleteProduct(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	exists, err := db.Exists(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.NoContent(http.StatusNotFound)
	}

	err = db.Delete(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
