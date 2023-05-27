package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

type Base struct {
	National_id string `json:"national_id" validate:"required"`
	Country     string `json:"country" validate:"required"`
	Type        string `json:"type" validate:"required"`
}

type ApiRequest struct {
	National_id    string `json:"national_id" validate:"required"`
	Country        string `json:"country" validate:"required"`
	Type           string `json:"type" validate:"required"`
	UserAuthorized bool   `json:"user_authorized" validate:"required"`
}

type CheckResponse struct {
	Base
	Check_id      string `json:"check_id"`
	Creation_date string `json:"creation_date"`
	Score         int    `json:"score"`
}

type Response struct {
	Body CheckResponse `json:"check"`
}

var API_URL = "https://api.checks.truora.com/v1"
var ErrBadRequest error = errors.New("bad request")

func BackgroundValidation(data io.ReadCloser) (*CheckResponse, error) {
	url := fmt.Sprintf("%s/checks", API_URL)
	dataValidated, err := validateData(data)
	if err != nil {
		log.Fatalf("Error validate %v", err.Error())
		return nil, err
	}
	response, err := createRequest("POST", url, dataValidated)
	if err != nil {
		log.Fatalf("Error response  %v", err.Error())
		return nil, err
	}
	return response, nil
}

func validateData(data io.ReadCloser) (*strings.Reader, error) {
	var request ApiRequest
	err := json.NewDecoder(data).Decode(&request)

	if err != nil {
		if err == io.EOF {
			log.Fatal("Error: JSON incompleto")
		} else {
			log.Fatal("Error al decodificar el JSON:", err)
		}
	}

	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		return nil, err
	}

	formData := url.Values{}
	formDataType := reflect.TypeOf(request)
	formDataValue := reflect.ValueOf(request)

	for i := 0; i < formDataType.NumField(); i++ {
		fieldName := formDataType.Field(i).Tag.Get("json")
		fieldValue := formDataValue.Field(i).Interface()
		fieldValueStr := fmt.Sprintf("%v", fieldValue)
		formData.Set(fieldName, fieldValueStr)
	}

	return strings.NewReader(formData.Encode()), nil
}

func ChecksDetails(checkId string) (*CheckResponse, error) {
	url := fmt.Sprintf("%s/checks/%s", API_URL, checkId)
	response, err := createRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error checks details %v", err.Error())
		return nil, err
	}
	return response, nil
}

func getApiKey() (string, error) {
	err := godotenv.Load(".env") //metodo para cargar nuestras variables de un file.
	if err != nil {
		log.Fatalf("Error loading .env file")
		return "", err
	}
	return os.Getenv("Truora_API_Key"), nil
}

func createRequest(method string, url string, body io.Reader) (*CheckResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	api_key, err := getApiKey()
	if err != nil {
		return nil, err
	}

	// Agregar encabezados a la solicitud
	req.Header.Add("Truora-API-Key", api_key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, ErrBadRequest
	}

	// Procesar la respuesta de la solicitud HTTP
	var response Response
	err = json.Unmarshal([]byte(bodyBytes), &response)

	if err != nil {
		return nil, err
	}
	return &response.Body, nil
}
