package model

type User struct {
	ID       string  `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Deposit  float32 `json:"deposit"`
	Role     string  `json:"role"`
}

type RegisterInput struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Deposit  float32 `json:"deposit"`
	Role     string  `json:"role"`
}

type LoginVar struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	ID           int     `json:"id"`
	ProductName  string  `json:"productName"`
	ProductPrice float32 `json:"productPrice"`
	SellerId     string  `json:"sellerId"`
}

type Results struct {
	TotalPrice       float32  `json:"totalPrice"`
	ProductPurchased string   `json:"productPurchased"`
	Change           []string `json:"change"`
}
