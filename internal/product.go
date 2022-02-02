package internal

import (
	"errors"
	"fmt"

	"github.com/realpamisa/model"
)

var products []model.Product

func CreateProduct(sellerId string, product model.Product) bool {
	if !FindProductByProductName(product.ProductName) {
		product.ID = len(products) + 1
		product.SellerId = sellerId
		products = append(products, product)
		return true
	}
	return false
}

func FindProductByProductName(productName string) bool {
	if len(products) > 0 {
		for _, u := range products {
			if u.ProductName == productName {
				return true
			}
		}
	}
	return false
}

func ViewAllProducts() []model.Product {
	return products
}

func UpdateProduct(productVar model.Product) bool {
	for _, p := range products {
		if p.ID == productVar.ID {
			p.ProductName = productVar.ProductName
			p.ProductPrice = productVar.ProductPrice
			return true
		}
	}
	return false
}

func GetProductByID(productID int) model.Product {
	for _, u := range products {
		if u.ID == productID {
			return u
		}
	}
	return model.Product{}
}

func BuyProduct(username string, productId int, amountOfProduct int) (*model.Results, error) {
	var (
		hundreds = 0
		fiftys   = 0
		tens     = 0
		twenties = 0
		fives    = 0
	)
	userRaw, err := GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	product := GetProductByID(productId)
	total := product.ProductPrice * float32(amountOfProduct)
	if userRaw.Deposit >= total {
		userRaw.Deposit = userRaw.Deposit - total
		change := userRaw.Deposit
		isSuccessUpdate := UpdateUser(username, userRaw)
		if !isSuccessUpdate {
			return nil, errors.New("Failed to update user")
		}
		if change/100 > 1 {
			hundreds = int(change) / 100
			change = change - float32(hundreds)*100
		}
		if change/50 > 1 {
			fiftys = int(change) / 50
			change = change - float32(fiftys)*50
		}
		if change/50 > 1 {
			fiftys = int(change) / 50
			change = change - float32(fiftys)*50
		}
		if change/20 > 1 {
			twenties = int(change) / 20
			change = change - float32(twenties)*20
		}
		if change/10 > 1 {
			tens = int(change) / 10
			change = change - float32(tens)*10
		}
		if change/5 > 1 {
			fives = int(change) / 5
			change = change - float32(fives)*5
		}
		changeArr := []string{fmt.Sprintf("%d x 100", hundreds), fmt.Sprintf("%d x 50", fiftys), fmt.Sprintf("%d x 20", twenties), fmt.Sprintf("%d x 10", tens), fmt.Sprintf("%d x 5", fives)}

		result := model.Results{
			TotalPrice:       total,
			ProductPurchased: product.ProductName,
			Change:           changeArr,
		}

		return &result, nil

	}
	return nil, errors.New("Not enough deposit money")
}

func DeleteProduct(id string) bool {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return true
		}
	}
	return false
}
