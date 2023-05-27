package controllers

import (
	"fmt"
	"net/http"

	"github.com/raisa320/Labora-wallet/repositories"
	"github.com/raisa320/Labora-wallet/repositories/postgres"
)

var walletRepository repositories.Wallet

func init() {
	// INYECCION DE DEPEDENCIA
	// base de datos (postgres)
	walletRepository = postgres.NewWalletStorage()
}

func GetAll(response http.ResponseWriter, request *http.Request) {
	wallets, err := walletRepository.GetAll()

	if err != nil {
		errMsg := fmt.Sprintf("error while retrieving wallets : '%v'", err)
		http.Error(response, errMsg, http.StatusInternalServerError)
		return
	}

	WriteJsonResponse(response, http.StatusOK, wallets)
}

func CreateWallet(response http.ResponseWriter, request *http.Request) {
	dataValidated, err := BackgroundValidation(request.Body)
	fmt.Println(err)
	WriteJsonResponse(response, http.StatusOK, dataValidated)
}
