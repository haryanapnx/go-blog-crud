package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	SuccessStatus = "success"
	SuccessCode   = "0"

	FailedStatus = "fail"
	FailCode     = "000002"

	ValidationStatus = "error validation"
	ValidationCode   = "000003"

	NotFoundStatus = "not found"
	NotFoundCode   = "000004"

	ErrorStatus = "error"
	ErrorCode   = "000005"

	DefaultError = "A connection attempt failed because connected host has failed to respond"
)

type ResponseStatus struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func handleStatus(statusCode int, message string, data interface{}) *ResponseStatus {
	response := &ResponseStatus{
		Code:   statusCode,
		Status: message,
		Data:   data,
	}
	return response
}

func JSON(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.WriteHeader(statusCode)
	// if statusCode < 400 {
	// 	message = SuccessStatus
	// }
	response := handleStatus(statusCode, message, data)
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, err.Error(), nil)
		return
	}
	JSON(w, http.StatusBadRequest, DefaultError, nil)
}
