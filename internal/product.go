package internal

import "github.com/realpamisa/model"

var products []model.Product

func CreateProduct(product model.Product) bool {
	if !FindProductByProductName(product.ProductName) {
		product.ID = len(products) + 1
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

func GetProductByProductName(productName string) model.Product {
	for _, u := range products {
		if u.ProductName == productName {
			return u
		}
	}
	return model.Product{}
}

func BuyProduct(username, productName string) bool {
	userRaw := GetUserByUsername(username)
	product := GetProductByProductName(productName)
	if userRaw.Deposit >= product.ProductPrice {
		userRaw.Deposit = userRaw.Deposit - product.ProductPrice
		isSuccessUpdate := UpdateUser(username, userRaw)
		if !isSuccessUpdate {
			return false
		}
		return true
	}
	return false
}
