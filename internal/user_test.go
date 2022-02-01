package internal

import (
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
	validUsername := "pamisa123"
	if !createUserTest(t) {
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
	if !createUserTest(t) {
		t.Errorf("got unexpected err creating user test")
	}

	usersTest := GetAllUsers()
	if len(usersTest) == 0 {
		t.Errorf("got unexpected err - got: %v , wanted: %v", len(usersTest), len(users))
	}
}

func createUserTest(t *testing.T) bool {
	validRequest := model.RegisterInput{
		Username: "pamisa123",
		Password: "123123",
		Role:     "buyer",
	}
	isSuccess := Register(validRequest)
	return isSuccess
}

func TestDeposit(t *testing.T) {
	validUsername := "pamisa123"
	if !createUserTest(t) {
		t.Errorf("got unexpected err creating user test")
	}
	validRequest := model.User{
		Deposit: 100,
	}
	invalidRequest := model.User{
		Deposit: 101,
	}
	tests := []struct {
		Name              string
		Username          string
		Request           model.User
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
			isSuccess := UpdateUser(tc.Username, tc.Request)
			if isSuccess != tc.ExpectedErrorBool {
				t.Errorf("got unexpected err - got: %v , wanted: %v", isSuccess, tc.ExpectedErrorBool)
			}
		})
	}

}
