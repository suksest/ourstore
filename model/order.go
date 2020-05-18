package model

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Order ...
type Order struct {
	ID        uint64
	ProductID uint64
	Quantity  uint64
	Customer  Customer
}

var (
	orders   = []Order{}
	orderseq = uint64(1)
)

//Create ...
func (o *Order) Create(c echo.Context) error {
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
	ord := new(Order)
	if err := c.Bind(ord); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid input",
		})
	}
	ord.ID = orderseq
	ord.ProductID = uint64(id)
	if products[uint64(id-1)].Quantity < ord.Quantity {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "out of stock",
		})
	}
	products[uint64(id-1)].Quantity -= ord.Quantity
	orderseq++
	return c.JSON(http.StatusOK, ord)
}
