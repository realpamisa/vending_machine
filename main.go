// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	jwt "github.com/golang-jwt/jwt"
// 	"github.com/gorilla/mux"
// 	"github.com/realpamisa/pkg/utils/token"
// )

// var mySigningKey = []byte("makodarklord")

// // func Register(w http.ResponseWriter, r *http.Request) {
// // 	raw := r.URL.Query()
// // 	username := raw.Get("username")
// // 	password := raw.Get("password")

// // 	validToken, err := token.New(username, password)
// // 	if err != nil {
// // 		fmt.Fprintf(w, err.Error())
// // 	}
// // 	fmt.Fprintf(w, validToken)
// // }

// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		if r.Header["Token"] != nil {
// 			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 					return nil, fmt.Errorf("There was an error")
// 				}
// 				return mySigningKey, nil
// 			})

// 			if err != nil {
// 				w.WriteHeader(http.StatusInternalServerError)
// 				fmt.Fprintf(w, "Not good")
// 			}
// 			if token.Valid {
// 				next.ServeHTTP(w, r)
// 			}
// 		} else {
// 			fmt.Fprintf(w, "Not Authorized")
// 		}
// 	})
// }

// // func isAuthorize(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
// // 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// // 		if r.Header["Token"] != nil {
// // 			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// // 				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// // 					return nil, fmt.Errorf("There was an error")
// // 				}
// // 				return mySigningKey, nil
// // 			})

// // 			if err != nil {
// // 				w.WriteHeader(http.StatusInternalServerError)
// // 				fmt.Fprintf(w, "Not good")
// // 			}
// // 			if token.Valid {
// // 				endpoint(w, r)
// // 			}
// // 		} else {
// // 			fmt.Fprintf(w, "Not Authorized")
// // 		}
// // 	})
// // }

// func Login(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Okay good")

// }

// // func DecodeToken(w http.ResponseWriter, token string) {

// // }

// func GenerateToken(username, password string) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["user"] = username
// 	claims["password"] = password

// 	tokenString, err := token.SignedString(mySigningKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }
// func handleRequest() {
// 	r := mux.NewRouter()
// 	r.HandleFunc("/register", Register).Methods("GET")
// 	r.HandleFunc("/login", loggingMiddleware(http.HandlerFunc(server.tagHandler))).Methods("POST")
// 	// http.HandleFunc("/register", Register)
// 	// http.Handle("/login", isAuthorize(Login))
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
// func main() {
// 	fmt.Println("Vending Machine")
// 	handleRequest()
// }
