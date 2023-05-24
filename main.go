package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/raisa320/Labora-wallet/config"
)

func myhandlers() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Inbit  api"))
	}).Methods("GET")

	// Configura las opciones de CORS. Por ejemplo, permite todas las origenes:
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	//allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	// Envolviendo tus rutas con CORS.
	handler := handlers.CORS(allowedOrigins, allowedMethods)(router)
	return handler
}

func main() {
	port := flag.String("port", ":9000", "The server port")
	flag.Parse()

	//services.InitDB()
	//services.Db.PingOrDie()

	router := myhandlers()

	server := config.NewServer(*port, router)

	log.Fatal(server.Start())
}
