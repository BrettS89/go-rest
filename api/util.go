package api

import (
	"encoding/json"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request) error
type HandlerWrapper func(w http.ResponseWriter, r *http.Request)

func WrapperHandler(requestHandler Handler) HandlerWrapper {
	return func(w http.ResponseWriter, r *http.Request) {
		err := requestHandler(w, r)

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(500)

			errorRespObj := map[string]string{
				"message": err.Error(),
			}

			json.NewEncoder(w).Encode(errorRespObj)
		}
	}
}

type HTTPResponse[T any] struct {
	StatusCode int
	Body       T
	Headers    []map[string]string
}
