package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/suksest/ourstore/model"
)

func main() {
	e := echo.New()
	p := new(model.Product)
	o := new(model.Order)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("products/:id", p.GetProductByID)
	e.GET("products", p.GetProducts)
	e.POST("products/:id/order", o.Create)
	e.Logger.Fatal(e.Start(":1323"))
}
