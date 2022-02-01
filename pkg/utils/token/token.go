package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("makodruglord")

type Claims struct {
	UserID   string `json:"userId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func New(username string, role string, userId string) (*string, error) {
	expirationTime := time.Now().Add((24 * 30) * time.Hour)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserID:   userId,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return nil, err
	}
	return &tokenString, nil

}

func Decode(bearerToken string) (*Claims, error) {
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//return &Claims{Email: fmt.Sprintf("%v", claims["email"]), Role: token.Claims.Role , Name: fmt.Sprintf("%v", claims["name"]), UserId: fmt.Sprintf("%v", claims["userId"])}, err
		return &Claims{Username: fmt.Sprintf("%v", claims["username"]), Role: fmt.Sprintf("%v", claims["role"]), UserID: fmt.Sprintf("%v", claims["userId"])}, err
	} else {
		return nil, err
	}
}
