package controllers

import (
	"github.com/raisa320/Labora-wallet/repositories"
	"github.com/raisa320/Labora-wallet/repositories/postgres"
)

var logRepository repositories.Log

func init() {
	// INYECCION DE DEPEDENCIA
	// base de datos (postgres)
	logRepository = postgres.NewLogStorage()
}
