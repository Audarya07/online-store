package app

import (
	"errors"
	"fmt"
)

type ProductServiceImpl struct {
	Products []Product
}

func (ps *ProductServiceImpl) GetAllProducts() ([]Product, error) {
	return ps.Products, nil
}

func (ps *ProductServiceImpl) GetProductByID(id int) (Product, error) {
	for _, product := range ps.Products {
		if product.ID == id {
			return product, nil
		}
	}
	return Product{}, errors.New("Product does not exist")
}

type CartServiceImpl struct {
	Carts          map[int]Cart
	ProductService *ProductServiceImpl
}

func (cs *CartServiceImpl) UpdateCart(userID int, productIDs []int) error {
	// Check if product IDs are valid and present in the product list
	for _, productID := range productIDs {
		_, err := cs.ProductService.GetProductByID(productID)
		if err != nil {
			errMsg := fmt.Sprintf("Product not found with id %d", productID)
			return errors.New(errMsg)
		}
	}

	cart, cartFound := cs.Carts[userID]
	if !cartFound {
		cart = Cart{UserID: userID, ProductIDs: []int{}}
	}

	// Add new product IDs to the cart
	for _, productID := range productIDs {
		found := false
		for _, existingProductID := range cart.ProductIDs {
			if productID == existingProductID {
				found = true
				break
			}
		}
		if !found {
			cart.ProductIDs = append(cart.ProductIDs, productID)
		}
	}

	// Remove product IDs from the cart that are not in the updated product IDs list
	updatedProductIDsMap := make(map[int]bool)
	for _, productID := range productIDs {
		updatedProductIDsMap[productID] = true
	}

	updatedProductIDs := []int{}
	for _, existingProductID := range cart.ProductIDs {
		if updatedProductIDsMap[existingProductID] {
			updatedProductIDs = append(updatedProductIDs, existingProductID)
		}
	}

	cart.ProductIDs = updatedProductIDs
	cs.Carts[userID] = cart

	return nil
}

func (cs *CartServiceImpl) GetCart(userID int) (Cart, error) {
	cart, ok := cs.Carts[userID]
	if !ok {
		errMsg := fmt.Sprintf("Cart not found for user id %d", userID)
		return Cart{}, errors.New(errMsg)
	}
	return cart, nil
}
