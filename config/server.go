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

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Init  api"))
	}).Methods("GET")

	router.HandleFunc("/api/v1/wallets", controllers.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/wallets", controllers.CreateWallet).Methods("POST")

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
