package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/raisa320/Labora-wallet/models"
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
	var data MyResponse = MyResponse{}
	wallets, err := walletRepository.GetAll()

	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
		return
	}

	WriteJsonResponse(response, http.StatusOK, wallets)
}

func CreateWallet(response http.ResponseWriter, request *http.Request) {
	var data MyResponse = MyResponse{}
	dataValidated, err := BackgroundValidation(request.Body)
	if err != nil {
		data.SetSimpleMessage(err.Message)
		WriteJsonResponse(response, err.StatusCode, data.Body)
		return
	}
	var log models.Log = models.Log{
		Person_id: dataValidated.National_id,
		Date:      dataValidated.Creation_date,
		Country:   dataValidated.Country,
		Check_id:  dataValidated.Check_id,
		Type:      "Wallet",
		Message:   "Ok",
	}
	fmt.Println("SCORE:", dataValidated.Score)
	log.SetStatus(dataValidated.Score)

	_, errCreate := logRepository.Create(request.Context(), log)
	if errCreate != nil {
		data.SetSimpleMessage(errCreate.Error())
		WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
		return
	}

	if log.Status {
		var wallet models.Wallet = models.Wallet{
			Person_id:  dataValidated.National_id,
			Date:       dataValidated.Creation_date,
			Country:    dataValidated.Country,
			PersonName: fmt.Sprintf("%s %s", dataValidated.Summary.NamesFound[0].FirstName, dataValidated.Summary.NamesFound[0].LastName),
		}
		_, err := walletRepository.Create(request.Context(), wallet)
		if err != nil {
			data.SetSimpleMessage(err.Error())
			WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
			return
		}
		data.StatusCode = http.StatusCreated
		data.Body = wallet
	} else {
		data.StatusCode = http.StatusOK
		data.SetSimpleMessage("Tu billetera no ha sido creada debido a que se encontraron antecedentes en tus registros")
	}

	WriteJsonResponse(response, data.StatusCode, data.Body)
}

func UpdateWallet(response http.ResponseWriter, request *http.Request) {
	var data MyResponse = MyResponse{}
	id, err := UrlParamInt(request, "id")
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusBadRequest, data.Body)
		return
	}

	var wallet models.Wallet
	json.NewDecoder(request.Body).Decode(&wallet)

	walletUpdated, err := walletRepository.Update(request.Context(), *id, wallet)
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
		return
	}

	if walletUpdated == nil {
		data.SetSimpleMessage("Wallet not found")
		WriteJsonResponse(response, http.StatusNotFound, data.Body)
		return
	}
	WriteJsonResponse(response, http.StatusOK, walletUpdated)
}

func DeleteWallet(response http.ResponseWriter, request *http.Request) {
	var data MyResponse = MyResponse{}
	id, err := UrlParamInt(request, "id")
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusBadRequest, data.Body)
		return
	}

	success, err := walletRepository.Delete(*id)
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
		return
	}
	if !success {
		data.SetSimpleMessage("Wallet not found")
		WriteJsonResponse(response, http.StatusNotFound, data.Body)
		return
	}
	data.SetSimpleMessage("Wallet deleted successfully")
	WriteJsonResponse(response, http.StatusOK, data.Body)
}

func StatusWallet(response http.ResponseWriter, request *http.Request) {
	var data MyResponse = MyResponse{}
	personId, err := UrlQueryParam(request, "personId")
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusBadRequest, data.Body)
		return
	}
	log, err := logRepository.GetByPersonId(*personId)
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
		return
	}
	if log == nil {
		message := fmt.Sprintf("No record was found of a virtual wallet associated with the personal identification number: %v", *personId)
		data.SetSimpleMessage(message)
		WriteJsonResponse(response, http.StatusNotFound, data.Body)
		return
	}
	jsonResponse := map[string]interface{}{
		"StatusWallet": log.GetStatus(),
	}
	WriteJsonResponse(response, http.StatusOK, jsonResponse)
}

func GetWalletById(response http.ResponseWriter, request *http.Request) {
	var data MyResponse = MyResponse{}
	id, err := UrlParamInt(request, "id")
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusBadRequest, data.Body)
		return
	}

	wallet, err := walletRepository.GetById(*id)
	if err != nil {
		data.SetSimpleMessage(err.Error())
		WriteJsonResponse(response, http.StatusInternalServerError, data.Body)
		return
	}

	if wallet == nil {
		data.SetSimpleMessage("Wallet not found")
		WriteJsonResponse(response, http.StatusNotFound, data.Body)
		return
	}
	WriteJsonResponse(response, http.StatusOK, wallet)
}
