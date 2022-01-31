package middleware

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/realpamisa/pkg/utils/token"
	"github.com/realpamisa/server/response"
)

var UserCtxKey = &contextKey{"userId"}
var ipAddressKey = &contextKey{"ipAddress"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the session into context
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			if len(splitToken) != 2 {
				w.Header().Set("Content-Type", "application/json")
				response.ERROR(w, http.StatusBadRequest, errors.New("Not bearer token"))
				return
			}
			reqToken = splitToken[1]
			fmt.Println(reqToken)
			claims, err := token.Decode(reqToken)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				response.ERROR(w, http.StatusBadRequest, err)
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), UserCtxKey, claims)
			//and call the next with our new context
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
func AdminMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			if len(splitToken) != 2 {
				w.Header().Set("Content-Type", "application/json")
				response.ERROR(w, http.StatusBadRequest, errors.New("Not bearer token"))
				return
			}
			reqToken = splitToken[1]
			claims, err := token.Decode(reqToken)

			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				response.ERROR(w, http.StatusBadRequest, err)
				return
			}
			for _,u := range claims.Role{
				if u.Role 
			}
			if claims.Role != "admin" {

				w.Header().Set("Content-Type", "application/json")
				response.ERROR(w, http.StatusBadRequest, errors.New("Not admin"))
				return
			}

			// put it in context
			ctx := context.WithValue(r.Context(), UserCtxKey, claims)
			//and call the next with our new context
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func GetClaims(r *http.Request) (*token.Claims, error) {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		return nil, errors.New("No token detected")
	}
	reqToken = splitToken[1]
	claims, err := token.Decode(reqToken)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
