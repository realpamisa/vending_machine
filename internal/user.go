package internal

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/realpamisa/model"
	"github.com/realpamisa/pkg/utils/token"
)

var (
	users []model.User
	role  = "buyer"
)

func Register(user model.RegisterInput) bool {
	if user.Role != "seller" {
		user.Role = role
	}
	if !FindUserByUsername(user.Username) {
		id := len(users) + 1
		newUser := model.User{
			ID:       fmt.Sprint(id),
			Username: user.Username,
			Password: user.Password,
			Deposit:  user.Deposit,
			Role:     user.Role,
			IsLogin:  false,
		}
		users = append(users, newUser)
		return true
	}

	return false
}

func Login(loginVar model.LoginVar) (*string, error) {
	user, err := GetUserByUsername(loginVar.Username)
	if err != nil {
		return nil, err
	}
	if user.Password == loginVar.Password {
		var validToken *string
		validToken, err := token.New(user.Username, user.Role, user.ID)
		if err != nil {
			return nil, err
		}
		user.IsLogin = true
		if !UpdateUser(loginVar.Username, user) {
			return nil, errors.New("Failed to update isLogin")
		}
		return validToken, nil
	}
	return nil, errors.New("invalid credentials")
}

func Logout(username string) bool {
	user, err := GetUserByUsername(username)
	if err != nil {
		return false
	}
	user.IsLogin = false
	if !UpdateUser(username, user) {
		return false
	}
	return true
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

func GetUserByUsername(username string) (model.User, error) {
	if len(users) > 0 {
		for _, u := range users {
			if u.Username == username {
				return u, nil
			}
		}
	}
	return model.User{}, errors.New("Invalid Username")
}

func GetAllUsers() []model.User {
	return users
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func UpdateUser(username string, user model.User) bool {

	for i, u := range users {
		if u.Username == username {
			users[i] = user
			return true
		}
	}
	return false
}

func Deposit(username string, money int) bool {
	if money == 5 || money == 10 || money == 20 || money == 50 || money == 100 {
		for _, u := range users {
			if u.Username == username {
				newUser := model.User{
					ID:       u.ID,
					Username: u.Username,
					Password: u.Password,
					Deposit:  u.Deposit + float32(money),
					Role:     u.Role,
				}
				isSuccess := UpdateUser(username, newUser)

				if !isSuccess {
					return false
				}
				return true
			}
		}
	}

	return false
}

func ResetDeposit(username string) bool {
	for _, u := range users {
		if u.Username == username {
			newUser := model.User{
				ID:       u.ID,
				Username: u.Username,
				Password: u.Password,
				Deposit:  0,
				Role:     u.Role,
			}
			isSuccess := UpdateUser(username, newUser)
			if !isSuccess {
				return false
			}
			return true
		}
	}
	return false
}
