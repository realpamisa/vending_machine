package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/realpamisa/pkg/utils/token"
	"github.com/realpamisa/server/response"
)

var UserCtxKey = &contextKey{"userId"}

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
func SellerMiddleware() func(http.Handler) http.Handler {
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
			if claims.Role != "seller" {

				w.Header().Set("Content-Type", "application/json")
				response.ERROR(w, http.StatusBadRequest, errors.New("Not seller"))
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

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func GetUserCtx(ctx context.Context) (*token.Claims, error) {
	raw, ok := ctx.Value(UserCtxKey).(*token.Claims)
	if ok {
		return raw, nil
	}
	return nil, errors.New("invalid token")

}
