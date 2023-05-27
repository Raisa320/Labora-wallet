package main

import (
	"flag"
	"log"

	"github.com/raisa320/Labora-wallet/config"
)

func main() {
	port := flag.String("port", ":9000", "The server port")
	flag.Parse()

	//services.InitDB()
	//services.Db.PingOrDie()

	server := config.NewServer(*port)

	log.Fatal(server.Start())
}
