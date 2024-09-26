package main

import (
	"encoding/json"
	"net/http"
	"reflect"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ UTILITY FUNCTIONS ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// makeHTTPHandlerFunc is a function that wraps an APIFunc in a http.HandlerFunc
func makeHTTPHandlerFunc(apiFunc APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := apiFunc(w, r)
		if err != nil {
			// Handle ERR
			writeJSON(w, err.Code, err.Error())
			return
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) *APIError {
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return &APIError{
			ErrorMessage: err.Error(),
			Code:         http.StatusInternalServerError,
		}
	}
	return nil
}

func getFields(v any) []interface{} {
	s := reflect.ValueOf(v).Elem()
	numFields := s.NumField()
	fields := make([]interface{}, numFields)

	for i := 0; i < numFields; i++ {
		field := s.Field(i)
		fields[i] = field.Addr().Interface()
	}

	return fields
}
