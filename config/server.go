package config

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/raisa320/Labora-wallet/controllers"
)

type Server struct {
	listenAddr string
	handler    http.Handler
}

func buildRouter() http.Handler {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/wallets", controllers.GetAll).Methods("GET")
	api.HandleFunc("/wallets/{id:[0-9]+}", controllers.GetWalletById).Methods("GET")
	api.HandleFunc("/wallets", controllers.CreateWallet).Methods("POST")
	api.HandleFunc("/wallets/{id:[0-9]+}", controllers.UpdateWallet).Methods("PUT")
	api.HandleFunc("/wallets/{id:[0-9]+}", controllers.DeleteWallet).Methods("DELETE")
	api.HandleFunc("/wallets/status", controllers.StatusWallet).Methods("GET")
	api.HandleFunc("/transactions/{id:[0-9]+}", controllers.GetTransaction).Methods("GET")
	api.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")

	// Configura las opciones de CORS. Por ejemplo, permite todas las origenes:
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	//allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	// Envolviendo tus rutas con CORS.
	handler := handlers.CORS(allowedOrigins, allowedMethods)(router)
	return handler
}

func NewServer(listenAddr string) *Server {
	router := buildRouter()
	return &Server{
		listenAddr: listenAddr,
		handler:    router,
	}
}

func (s *Server) Start() error {
	fmt.Println("Server running on port", s.listenAddr)
	return http.ListenAndServe(s.listenAddr, s.handler)
}
