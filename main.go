package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/raisa320/API/config"
	"github.com/raisa320/API/controllers"
	"github.com/raisa320/API/services"
)

func myhandlers() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/items", controllers.GetItems).Methods("GET")
	router.HandleFunc("/itemsPerPage", controllers.GetItemsPage).Methods("GET")
	router.HandleFunc("/items", controllers.SaveItem).Methods("POST")
	router.HandleFunc("/items/{id}", controllers.GetItem).Methods("GET")
	router.HandleFunc("/item", controllers.SearchItemByCustomer).Methods("GET")
	router.HandleFunc("/items/{id}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", controllers.DeleteItem).Methods("DELETE")

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

	services.InitDB()
	services.Db.PingOrDie()

	router := myhandlers()

	server := config.NewServer(*port, router)

	log.Fatal(server.Start())
}
