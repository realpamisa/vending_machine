package internal

import (
	"encoding/json"
	"errors"

	"github.com/realpamisa/model"
	"github.com/realpamisa/pkg/utils/token"
)

var (
	users []model.User
	role  = "user"
)

func Register(user model.User) bool {
	if user.Role != "admin" {
		user.Role = role
	}
	if !FindUserByUsername(user.Username) {
		user.ID = len(users) + 1
		users = append(users, user)
		return true
	}
	return false
}

func Login(loginVar model.LoginVar) (*string, error) {
	user := GetUserByUsername(loginVar.Username)
	if user.Password == loginVar.Password {

		var validToken *string
		validToken, err := token.New(user.Username, user.Role)
		if err != nil {
			return nil, err
		}
		return validToken, nil
	}
	return nil, errors.New("invalid credentials")

}

func FindUserByUsername(username string) bool {

	if len(users) > 0 {
		for _, u := range users {
			if u.Username == username {
				return true
			}
		}
	}
	return false
}

func GetUserByUsername(username string) model.User {
	if len(users) > 0 {
		for _, u := range users {
			if u.Username == username {
				return u
			}
		}
	}
	return model.User{}
}

func GetAllUsers() []model.User {
	return users
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func UpdateUser(username string, user model.User) bool {

	for _, u := range users {
		if u.Username == username {
			if user.Deposit != 0 {
				u.Deposit = user.Deposit
			}
			if user.Username != "" {
				u.Username = user.Username
			}
			return true
		}
	}
	return false
}

func Deposit(username string, money int) bool {
	if money == 5 || money == 10 || money == 20 || money == 50 || money == 100 {
		for _, u := range users {
			if u.Username == username {
				u.Deposit = u.Deposit + money
				return true
			}
		}
	}

	return false
}

func ResetDeposit(username string) bool {
	for _, u := range users {
		if u.Username == username {
			u.Deposit = 0
			return true
		}
	}
	return false
}
