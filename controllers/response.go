package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Error struct {
	StatusCode int
	Message    string
}

type MyResponse struct {
	StatusCode int
	Body       interface{}
}

func (response *MyResponse) SetSimpleMessage(message string) {
	response.Body = map[string]interface{}{
		"Message": message,
	}
}

func WriteJsonResponse(response http.ResponseWriter, status int, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error while mashalling object %v, trace: %+v", data, err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	_, err = response.Write(bytes)
	if err != nil {
		fmt.Printf("error while writting bytes to response writer: %+v", err)
	}
}

func UrlQueryParam(request *http.Request, name string) (*string, error) {
	param := request.URL.Query().Get(name)
	if len(param) == 0 {
		return nil, fmt.Errorf("no url param present with the name: %v", name)
	}
	return &param, nil
}

func UrlParamInt(request *http.Request, name string) (*int, error) {
	param := mux.Vars(request)[name]
	paramInt, err := strconv.Atoi(param)
	if err != nil {
		return nil, fmt.Errorf("can not parse id from %v", param)
	}
	return &paramInt, nil
}
