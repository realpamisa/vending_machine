package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/realpamisa/internal"
	"github.com/realpamisa/middleware"
	"github.com/realpamisa/model"
	"github.com/realpamisa/server/response"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var registerInput model.RegisterInput
	values := r.URL.Query()
	if values.Get("username") != "" {
		registerInput.Username = values.Get("username")
	}
	if values.Get("password") != "" {
		registerInput.Password = values.Get("password")
	}
	if values.Get("role") != "" {
		registerInput.Role = values.Get("role")
	}

	isSuccess := internal.Register(registerInput)

	if !isSuccess {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Failed to create account"))
		return
	} else {
		response.JSON(w, http.StatusCreated, isSuccess)
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	var loginVar model.LoginVar

	values := r.URL.Query()
	if values.Get("username") != "" {
		loginVar.Username = values.Get("username")
	}
	if values.Get("password") != "" {
		loginVar.Password = values.Get("password")
	}

	token, err := internal.Login(loginVar)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	response.JSON(w, http.StatusCreated, token)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r)
	if err != nil {
		response.ERROR(w, 400, err)
		return
	}
	if !internal.Logout(claims.Username) {
		response.ERROR(w, 400, err)
		return
	}
	response.JSON(w, 200, true)
	return
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	data := internal.GetAllUsers()
	if len(data) != 0 {
		response.JSON(w, 200, data)
		return
	}
	response.ERROR(w, http.StatusUnprocessableEntity, errors.New("No User"))
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r)
	if err != nil {
		response.ERROR(w, 400, err)
		return
	}
	data, err := internal.GetUserByUsername(claims.Username)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	response.JSON(w, 200, data)
	return
}
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productVar model.Product

	values := r.URL.Query()
	if values.Get("productPrice") != "" {
		productPrice, err := strconv.Atoi(values.Get("productPrice"))
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error parsing productPrice"))
			return
		}
		productVar.ProductPrice = float32(productPrice)
	}
	if values.Get("productName") != "" {
		productVar.ProductName = values.Get("productName")
	}
	claims, err := middleware.GetClaims(r)
	if err != nil {
		response.ERROR(w, 400, err)
		return
	}
	isSuccess := internal.CreateProduct(claims.UserID, productVar)

	if !isSuccess {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Failed to create product"))
		return
	} else {
		response.JSON(w, http.StatusCreated, isSuccess)
		return
	}
}

func ViewAllProducts(w http.ResponseWriter, r *http.Request) {
	products := internal.ViewAllProducts()
	if len(products) != 0 {
		response.JSON(w, 200, products)
		return
	}
	response.ERROR(w, http.StatusUnprocessableEntity, errors.New("No Product"))
	return
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var productVar model.Product

	values := r.URL.Query()
	if values.Get("productPrice") != "" {
		productPrice, err := strconv.Atoi(values.Get("productPrice"))
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error parsing productPrice"))
		}
		productVar.ProductPrice = float32(productPrice)
	}
	if values.Get("productName") != "" {
		productVar.ProductName = values.Get("productName")
	}
	if values.Get("ID") != "" {
		id, err := strconv.Atoi(values.Get("id"))
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error parsing id"))
			return
		}
		productVar.ID = id
	}
	response.JSON(w, 200, true)
	return
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	if !internal.DeleteProduct(values.Get("productId")) {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error delete user"))
	}
}

func Deposit(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	claims, err := middleware.GetClaims(r)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error parsing claims"))
		return
	}
	if values.Get("deposit") != "" {
		deposit, err := strconv.Atoi(values.Get("deposit"))
		if err != nil {
			response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error parsing deposit"))
			return
		}

		isDepositSuccess := internal.Deposit(claims.Username, deposit)
		if !isDepositSuccess {
			response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error money"))
			return
		}
		response.JSON(w, 200, true)
		return
	}
	response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error deposit money"))
	return
}

func ResetDeposit(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.GetClaims(r)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error parsing claims"))
		return
	}
	isResetSuccess := internal.ResetDeposit(claims.Username)
	if !isResetSuccess {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error Reset Deposit"))
		return
	}
	response.JSON(w, 200, true)
	return
}

func BuyProduct(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	claims, err := middleware.GetClaims(r)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error parsing claims"))
		return
	}
	productId, err := strconv.Atoi(values.Get("productId"))
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error parsing deposit"))
		return
	}
	amountOfProduct, err := strconv.Atoi(values.Get("amountOfProduct"))
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error parsing amountOfProduct"))
		return
	}
	data, err := internal.BuyProduct(claims.Username, productId, amountOfProduct)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	response.JSON(w, 200, data)
	return
}
