package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/raisa320/Labora-wallet/models"
	"github.com/raisa320/Labora-wallet/repositories"
	"github.com/raisa320/Labora-wallet/repositories/postgres"
)

var TransactionRepository repositories.Transaction

func init() {
	// INYECCION DE DEPEDENCIA
	// base de datos (postgres)
	TransactionRepository = postgres.NewTransactionStorage()
}

func GetTransaction(response http.ResponseWriter, request *http.Request) {
	var data MyResponse = MyResponse{}
	id, err := UrlParamInt(request, "id")
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusBadRequest, data.Body)
		return
	}
	transaction, err := TransactionRepository.GetById(request.Context(), *id)
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
		return
	}

	if transaction == nil {
		data.SetSimpleMessage("Transaction not found")
		WriteJsonResponse(response, http.StatusNotFound, data.Body)
		return
	}
	WriteJsonResponse(response, http.StatusOK, transaction)
}

func GetTransactionsByWallet(response http.ResponseWriter, request *http.Request) {
	var data MyResponse = MyResponse{}
	id, err := UrlParamInt(request, "id")
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusBadRequest, data.Body)
		return
	}
	transactions, err := TransactionRepository.GetByWallet(*id)
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
		return
	}

	if transactions == nil {
		data.SetSimpleMessage("Wallet id invalid")
		WriteJsonResponse(response, http.StatusNotFound, data.Body)
		return
	}
	WriteJsonResponse(response, http.StatusOK, transactions)
}

func CreateTransaction(response http.ResponseWriter, request *http.Request) {
	var data MyResponse = MyResponse{}
	var transaction models.Transaction
	json.NewDecoder(request.Body).Decode(&transaction)
	err := TransactionRepository.Create(request.Context(), transaction)
	if err != nil {
		var log models.Log = models.Log{
			Date:    time.Now(),
			Status:  false,
			Type:    "Transaction",
			Message: err.Error(),
		}
		logRepository.Create(request.Context(), log)
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
		return
	}
	data.SetSimpleMessage("Succesful transaction")
	WriteJsonResponse(response, http.StatusOK, data.Body)
}
