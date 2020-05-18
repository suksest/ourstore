package model

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo"
)

//Product represent single product
type Product struct {
	ID       uint64
	Name     string
	Quantity uint64
	Price    float64
}

//SafeProduct ...
type SafeProduct struct {
	Product Product
	mux     sync.Mutex
}

var (
	products = []Product{}
	seq      = uint64(1)
)

//GetProductByID return product with given ID
func (p *Product) GetProductByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}
	if len(products) < id {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}
	return c.JSON(http.StatusOK, products[uint64(id-1)])
}

//GetProducts return all products
func (p *Product) GetProducts(c echo.Context) error {
	if len(products) == 0 {
		products = append(products, Product{
			ID:       seq,
			Name:     "Sabun",
			Price:    500.0,
			Quantity: 3,
		})
	}
	return c.JSON(http.StatusOK, products)
}
