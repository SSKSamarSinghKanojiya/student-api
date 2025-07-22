package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string `json:"custom_status"`
	Error  string `json:"error"`
}

const (
	StatusOK     = "OK"
	StatsusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GenralError(err error) Response {
	return Response{
		Status: StatsusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) Response {
	var errMgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errMgs = append(errMgs, fmt.Sprintf("field %s is required field ", err.Field()))

		default:
			errMgs = append(errMgs, fmt.Sprintf("field %s is invalid", err.Field()))
		}
	}
	return Response{
		Status: StatsusError,
		Error: strings.Join(errMgs,", "),
	}
}
