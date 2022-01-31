package model

type User struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Deposit  int      `json:"deposit"`
	Role     []string `json:"role"`
}

type RegisterInput struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Deposit  int      `json:"deposit"`
	Role     []string `json:"role"`
}

type LoginVar struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Product struct {
	ID           int    `json:"id"`
	ProductName  string `json:"productName"`
	ProductPrice int    `json:"productPrice"`
}
