package main

import (
	"igloo/app"

	"github.com/labstack/echo/v4"
)

var products []app.Product

func main() {
	// Initialize mock data
	products = []app.Product{
		{ID: 1, Name: "Product 1", Description: "Description 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Description: "Description 2", Price: 100.25},
		{ID: 3, Name: "Product 3", Description: "Description 3", Price: 50.00},
		{ID: 4, Name: "Product 4", Description: "Description 4", Price: 1000},
		{ID: 5, Name: "Product 5", Description: "Description 5", Price: 76.67},
	}

	var carts = map[int]app.Cart{
		1: {UserID: 1, ProductIDs: []int{1, 2}},
	}

	// Initialize Echo instance
	e := echo.New()

	// Initialize services
	productService := &app.ProductServiceImpl{Products: products}
	cartService := &app.CartServiceImpl{Carts: carts, ProductService: productService}

	// Initialize API handler with services
	handler := &app.APIHandler{
		ProductService: productService,
		CartService:    cartService,
	}

	// Register API routes
	app.RegisterRoutes(e, handler)

	// Start the server
	e.Logger.Fatal(e.Start(":8000"))
}
