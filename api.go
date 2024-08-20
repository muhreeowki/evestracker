package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	store      Storage
	listenAddr string
}

type APIFunc func(http.ResponseWriter, *http.Request) *APIError

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	r := chi.NewRouter()
	r.Get("/", makeHTTPHandlerFunc(s.handleGetMidwife))

	log.Printf("EvesTracker API is running on port: %s", s.listenAddr)

	http.ListenAndServe(s.listenAddr, r)
}

func (s *APIServer) handleGetMidwife(w http.ResponseWriter, r *http.Request) *APIError {
	writeJSON(w, http.StatusOK, "Hello World")
	return nil
}

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

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
