package app

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type APIHandler struct {
	ProductService ProductService
	CartService    CartService
}

func (h *APIHandler) GetAllProducts(c echo.Context) error {
	products, err := h.ProductService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)
}

func (h *APIHandler) GetProductByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}
	product, err := h.ProductService.GetProductByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, "Product does not exist")
	}
	return c.JSON(http.StatusOK, product)
}

func (h *APIHandler) UpdateCart(c echo.Context) error {
	var cart Cart
	if err := c.Bind(&cart); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request body")
	}
	userID := cart.UserID
	err := h.CartService.UpdateCart(userID, cart.ProductIDs)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Product(s) added to cart", "cart": cart})
}

func (h *APIHandler) GetCart(c echo.Context) error {
	userID := 1 // Currently hardcoded for just 1 cart, can be used for cartId/userId coming from request
	cart, err := h.CartService.GetCart(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, cart)
}
