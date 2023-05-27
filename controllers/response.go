package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
