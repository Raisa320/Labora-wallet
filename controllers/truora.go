package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

type ApiRequest struct {
	National_id    string `json:"national_id" validate:"required"`
	Country        string `json:"country" validate:"required"`
	Type           string `json:"type" validate:"required"`
	UserAuthorized bool   `json:"user_authorized" validate:"required"`
}

type CheckResponse struct {
	National_id   string    `json:"national_id"`
	Country       string    `json:"country"`
	Type          string    `json:"type"`
	Check_id      string    `json:"check_id"`
	Creation_date time.Time `json:"creation_date"`
	Score         int       `json:"score"`
	Summary       struct {
		NamesFound []struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		} `json:"names_found"`
	} `json:"summary"`
}

type Response struct {
	Body CheckResponse `json:"check"`
}

var API_URL = "https://api.checks.truora.com/v1"

func BackgroundValidation(data io.ReadCloser) (*CheckResponse, *Error) {
	url := fmt.Sprintf("%s/checks", API_URL)
	dataValidated, err := validateData(data)
	if err != nil {
		return nil, err
	}

	response, err := createRequest("POST", url, dataValidated)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func validateData(data io.ReadCloser) (*strings.Reader, *Error) {
	var request ApiRequest
	err := json.NewDecoder(data).Decode(&request)

	if err != nil {
		var errResp Error
		if err == io.EOF {
			errResp = Error{StatusCode: http.StatusBadRequest, Message: "Error: JSON incompleto"}
		} else {
			errResp = Error{StatusCode: http.StatusInternalServerError, Message: fmt.Sprintf("Error al decodificar el JSON: %v", err)}
		}
		return nil, &errResp
	}

	validate := validator.New()

	err = validate.Struct(request)
	if err != nil {
		errResp := Error{StatusCode: http.StatusBadRequest, Message: err.Error()}
		return nil, &errResp
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

func ChecksDetails(checkId string) (*CheckResponse, *Error) {
	url := fmt.Sprintf("%s/checks/%s", API_URL, checkId)
	response, err := createRequest("GET", url, nil)
	if err != nil {
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

func createRequest(method string, url string, body io.Reader) (*CheckResponse, *Error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		errResp := Error{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		return nil, &errResp
	}

	api_key, err := getApiKey()
	if err != nil {
		errResp := Error{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		return nil, &errResp
	}

	// Agregar encabezados a la solicitud
	req.Header.Add("Truora-API-Key", api_key)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		errResp := Error{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		return nil, &errResp
	}

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		errResp := Error{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		return nil, &errResp
	}

	if resp.StatusCode >= 400 {
		errResp := Error{StatusCode: resp.StatusCode, Message: "Error response API"}
		return nil, &errResp
	}

	// Procesar la respuesta de la solicitud HTTP
	var response Response
	err = json.Unmarshal([]byte(bodyBytes), &response)

	if err != nil {
		errResp := Error{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		return nil, &errResp
	}
	return &response.Body, nil
}
