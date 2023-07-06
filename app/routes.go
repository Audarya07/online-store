package app

import (
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, handler *APIHandler) {
	e.GET("/products", handler.GetAllProducts)
	e.GET("/products/:id", handler.GetProductByID)
	e.PUT("/cart", handler.UpdateCart)
	e.GET("/cart", handler.GetCart)
}
