package app

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Cart struct {
	UserID     int   `json:"userId"`
	ProductIDs []int `json:"productIds"`
}

type ProductService interface {
	GetAllProducts() ([]Product, error)
	GetProductByID(id int) (Product, error)
}

type CartService interface {
	UpdateCart(userID int, productIDs []int) error
	GetCart(userId int) (Cart, error)
}
