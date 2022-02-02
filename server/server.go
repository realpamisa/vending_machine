package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/realpamisa/internal/handler"
	requestMiddleWare "github.com/realpamisa/middleware"
	"github.com/rs/cors"
)

type Server struct{}

func (srv Server) Start() {

	r := chi.NewRouter()
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"X-Requested-With", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Debug:            false,
		AllowCredentials: true,
	})
	r.Use(middleware.Timeout(60 * time.Second)) // 1minute timeout
	r.Use(cors.Handler)
	// r.Use(requestMiddleWare.Middleware())
	r.Get("/register", handler.Register)
	r.Post("/login", handler.Login)
	r.Mount("/seller", sellerRouter())
	r.Mount("/buyer", buyerRouter())
	r.Mount("/logout", logoutRouter())

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
	return
}

func logoutRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(requestMiddleWare.SellerMiddleware())
	r.Get("/", handler.Logout)
}

func sellerRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(requestMiddleWare.SellerMiddleware())
	r.Get("/", handler.GetUser)
	r.Get("/users", handler.GetUsers)
	r.Get("/product/create", handler.CreateProduct)
	r.Get("/product/all", handler.ViewAllProducts)
	r.Get("/product/update", handler.UpdateProduct)
	r.Get("/product/delete", handler.DeleteProduct)
	// r.Get("/accounts", adminListAccounts)
	return r
}

func buyerRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(requestMiddleWare.Middleware())
	r.Get("/", handler.GetUser)
	r.Get("/deposit", handler.Deposit)
	r.Get("/reset", handler.ResetDeposit)
	r.Get("/buy", handler.BuyProduct)
	// r.Get("/accounts", adminListAccounts)
	return r
}
