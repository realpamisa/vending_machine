package internal

import (
	"strings"
	"testing"

	"github.com/realpamisa/model"
)

func TestCreate(t *testing.T) {
	validRequest := model.RegisterInput{
		Username: "pamisa123",
		Password: "123123",
		Role:     "buyer",
	}
	invalidRequest := model.RegisterInput{
		Username: "pamisa123",
		Role:     "buyer",
	}
	tests := []struct {
		Name              string
		Request           model.RegisterInput
		ExpectedErrorBool bool
	}{
		{
			Name:              "ValidRegisterUser",
			Request:           validRequest,
			ExpectedErrorBool: true,
		},
		{
			Name:              "InvalidRegisterUser",
			Request:           invalidRequest,
			ExpectedErrorBool: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			isSuccess := Register(tc.Request)
			if isSuccess != tc.ExpectedErrorBool {
				t.Errorf("got unexpected err - got: %v , wanted: %v", isSuccess, tc.ExpectedErrorBool)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	validUsername := "pamisa1234"
	if !createUserTest(validUsername, t) {
		t.Errorf("got unexpected err creating user test")
	}
	invalidUsername := "zxcczxczx"
	Request := model.User{
		Username: "pamisa12322",
		Password: "123123",
	}

	tests := []struct {
		Name              string
		Username          string
		Request           model.User
		ExpectedErrorBool bool
	}{
		{
			Name:              "ValidUpdateUser",
			Username:          validUsername,
			Request:           Request,
			ExpectedErrorBool: true,
		},
		{
			Name:              "InvalidUpdateUser",
			Username:          invalidUsername,
			Request:           Request,
			ExpectedErrorBool: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			isSuccess := UpdateUser(tc.Username, tc.Request)
			if isSuccess != tc.ExpectedErrorBool {
				t.Errorf("got unexpected err - got: %v , wanted: %v", isSuccess, tc.ExpectedErrorBool)
			}
		})
	}
}

func TestViewUser(t *testing.T) {
	validUsername := "pamisa12345"
	if !createUserTest(validUsername, t) {
		t.Errorf("got unexpected err creating user test")
	}

	usersTest := GetAllUsers()
	if len(usersTest) == 0 {
		t.Errorf("got unexpected err - got: %v , wanted: %v", len(usersTest), len(users))
	}
}

func createUserTest(username string, t *testing.T) bool {
	validRequest := model.RegisterInput{
		Username: username,
		Password: "123123",
		Role:     "buyer",
		Deposit:  100,
	}
	isSuccess := Register(validRequest)
	return isSuccess
}

func TestDeposit(t *testing.T) {
	validUsername := "pamisa123456"
	if !createUserTest(validUsername, t) {
		t.Errorf("got unexpected err creating user test")
	}
	validRequest := 100
	invalidRequest := 101
	tests := []struct {
		Name              string
		Username          string
		Request           int
		ExpectedErrorBool bool
	}{
		{
			Name:              "ValidDeposit",
			Username:          validUsername,
			Request:           validRequest,
			ExpectedErrorBool: true,
		},
		{
			Name:              "InvalidDeposit",
			Username:          validUsername,
			Request:           invalidRequest,
			ExpectedErrorBool: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			isSuccess := Deposit(tc.Username, tc.Request)
			if isSuccess != tc.ExpectedErrorBool {
				t.Errorf("got unexpected err - got: %v , wanted: %v", isSuccess, tc.ExpectedErrorBool)
			}
		})
	}

}

func TestDeleteProduct(t *testing.T) {
	validUsername := "pamisa123456789"
	if !createUserTest(validUsername, t) {
		t.Errorf("got unexpected err creating user test")
	}
	if !createProductTest(validUsername, t) {
		t.Errorf("got unexpected err creating product test")
	}
	validRequest := "1"
	invalidRequest := "2"
	tests := []struct {
		Name              string
		Request           string
		ExpectedErrorBool bool
	}{
		{
			Name:              "ValidDeleteProduct",
			Request:           validRequest,
			ExpectedErrorBool: true,
		},
		{
			Name:              "InvalidDeleteProduct",
			Request:           invalidRequest,
			ExpectedErrorBool: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			isSuccess := DeleteProduct(tc.Request)
			if isSuccess != tc.ExpectedErrorBool {
				t.Errorf("got unexpected err - got: %v , wanted: %v", isSuccess, tc.ExpectedErrorBool)
			}
		})
	}

}

type BuyProductTest struct {
	Username        string
	ProductID       int
	AmountOfProduct int
}

func TestBuy(t *testing.T) {
	validUsername := "pamisa1234567"
	if !createUserTest(validUsername, t) {
		t.Errorf("got unexpected err creating user test")
	}
	if !createProductTest(validUsername, t) {
		t.Errorf("got unexpected err creating product test")
	}
	validRequest := BuyProductTest{
		Username:        validUsername,
		ProductID:       1,
		AmountOfProduct: 2,
	}
	invalidAmountOfMoney := BuyProductTest{
		Username:        validUsername,
		ProductID:       1,
		AmountOfProduct: 10,
	}
	invalidUsername := BuyProductTest{
		Username:        "invalid",
		ProductID:       1,
		AmountOfProduct: 2,
	}
	tests := []struct {
		Name              string
		Request           BuyProductTest
		ExpectedErrString string
	}{
		{
			Name:              "ValidDeposit",
			Request:           validRequest,
			ExpectedErrString: "",
		},
		{
			Name:              "InvalidAmountOfMoney",
			Request:           invalidAmountOfMoney,
			ExpectedErrString: "Not enough deposit money",
		},
		{
			Name:              "invalidUsername",
			Request:           invalidUsername,
			ExpectedErrString: "Invalid Username",
		},
	}
	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			_, err := BuyProduct(tc.Request.Username, tc.Request.ProductID, tc.Request.AmountOfProduct)
			if tc.ExpectedErrString != "" && err != nil && !strings.Contains(err.Error(), tc.ExpectedErrString) {
				t.Errorf("got unexpected err - got: %v , wanted: %v", err, tc.ExpectedErrString)
			}
		})
	}

}

func createProductTest(username string, t *testing.T) bool {
	validRequest := model.Product{
		ProductName:  "mako",
		ProductPrice: 20,
	}
	isSuccess := CreateProduct(username, validRequest)
	return isSuccess
}

func createUserSellerTest(t *testing.T) bool {
	validRequest := model.RegisterInput{
		Username: "pamisa123456789",
		Password: "123123",
		Role:     "seller",
	}
	isSuccess := Register(validRequest)
	return isSuccess
}
