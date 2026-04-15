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
	res, err := db.ListProduct()
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
	id, err := db.CreateProduct(product)
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

	exists, err := db.ExistsProduct(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.NoContent(http.StatusNotFound)
	}

	res, err := db.ReadProduct(id)
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

	exists, err := db.ExistsProduct(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.NoContent(http.StatusNotFound)
	}

	product, err := db.ReadProduct(id)
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

	err = db.UpdateProduct(id, product)
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

	exists, err := db.ExistsProduct(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.NoContent(http.StatusNotFound)
	}

	err = db.DeleteProduct(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// POST /carts
func CreateCart(c *echo.Context) error {
	id, err := db.CreateCart()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	c.Response().Header().Set("Location", "/carts/"+id.String())
	return c.NoContent(http.StatusCreated)
}

// GET /carts/:id
func ReadCart(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	exists, err := db.ExistsCart(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.NoContent(http.StatusNotFound)
	}

	res, err := db.ReadCart(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

// DELETE /carts/:id
func DeleteCart(c *echo.Context) error {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	exists, err := db.ExistsCart(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.NoContent(http.StatusNotFound)
	}

	err = db.DeleteCart(id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

// POST /carts/:id
func CreateCartItem(c *echo.Context) error {
	cidStr := c.Param("id")
	cid, err := uuid.Parse(cidStr)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	pidStr := c.FormValue("id")
	if pidStr == "" {
		return c.String(http.StatusBadRequest, "required value missing")
	}

	pid, err := uuid.Parse(pidStr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	exists, err := db.ExistsProduct(pid)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if !exists {
		return c.String(http.StatusBadRequest, "product "+pid.String()+" not found")
	}

	err = db.AddToCart(cid, pid)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusAccepted)
}
