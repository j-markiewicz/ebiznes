package main

import (
	"crypto/rand"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestLogger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("/products", listProducts)
	e.GET("/products/:id", readProduct)

	e.POST("/login", login)
	e.POST("/signup", signup)

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

type User struct {
	Username string `json:"username"`
	Method   string `json:"method"`
	Token    string `json:"token"`
}

var users = map[string]User{
	"test":     {"test", "password", "test"},
	"username": {"username", "password", "password"},
}

var tokens = map[string]string{}

// GET /products
func listProducts(c *echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

// GET /products/:id
func readProduct(c *echo.Context) error {
	return c.JSON(http.StatusOK, products[c.Param("id")])
}

// POST /login
func login(c *echo.Context) error {
	username := c.FormValue("username")
	method := c.FormValueOr("provider", "password")

	if users[username].Username != username {
		return c.String(http.StatusForbidden, "invalid username")
	}

	if users[username].Method != method {
		return c.String(http.StatusForbidden, "invalid authentication provider")
	}

	if method == "password" {
		password := c.FormValue("password")

		if users[username].Token != password {
			return c.String(http.StatusForbidden, "invalid password")
		}
	} else {
		return c.NoContent(http.StatusForbidden)
	}

	token := rand.Text()
	tokens[token] = username
	c.SetCookie(&http.Cookie{Name: "auth"})
	return c.NoContent(http.StatusOK)
}

// POST /signup
func signup(c *echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if users[username].Username != "" {
		return c.String(http.StatusForbidden, "invalid username")
	}

	if users[username].Username != "" {
		return c.String(http.StatusForbidden, "user already exists")
	}

	users[username] = User{
		Username: username,
		Method:   "password",
		Token:    password,
	}

	return c.NoContent(http.StatusOK)
}
