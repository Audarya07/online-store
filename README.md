# Online Store API

This repository contains the implementation of an API for an online store. The API provides functionality to retrieve the list of products, retrieve details of a product, add/remove product(s) to cart, and retrieve a user's cart.

## Problem Statement

Design an API for an online store with the following functionality:  
1. Retrieve list of all products.
2. Retrieve details of a product.
3. Add/remove product(s) to cart.
4. Retrieve a user's cart.

## API Description

### Retrieve List of All Products

- **Path:** `/products`
- **Method:** GET
- **Description:** Retrieves the list of all products available in the store.
- **Response:**
  ```json
  [
    {
      "id": 1,
      "name": "Product 1",
      "description": "Description 1",
      "price": 10.99
    },
    {
      "id": 2,
      "name": "Product 2",
      "description": "Description 2",
      "price": 19.99
    },
    ...
  ]
  ```

### Retrieve Details of a Product

  - **Path:** : `/products/{id}`
  - **Method**: GET
  - **Description:** Retrieves the details of a specific product.
  - **Response**:
  ```json
  {
  "id": 1,
  "name": "Product 1",
  "description": "Description 1",
  "price": 10.99
 } 
```
   
  
### Update Cart (Add/Remove Product)

  - **Path:** : `/cart`
  - **Method**: PUT
  - **Description:** Updates the cart. Update can be adition or removal of a product.
  - **Body**:
  ```json
  {
    "userId": 1,
    "productIds": [
        1,
        2,
        3
    ]
}
``` 

  - **Response**:

  ```json
  {
    "cart": {
        "userId": 1,
        "productIds": [
            1,
            2,
            3
        ]
    },
    "message": "Product(s) added to cart"
  }
  ```


### Retrieve User's Cart 

  - **Path:** : `/cart`
  - **Method**: GET
  - **Description:** Retrieves the user's cart.
  - **Response**:

  ```json
  {
    "userId": 1,
    "productIds": [
        1,
        2,
        3
    ]
}
  ```

## Setup

  - Clone the repository:
```shell 
git clone https://github.com/Audarya07/online-store.git
```

- Install the dependencies:

```shell
go get ./...
```

- Start the server:

```shell
    go run main.go
```

The server will be accessible at http://localhost:8000.

Feel free to explore and use the API endpoints as described above. 
