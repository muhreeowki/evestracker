package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	r.Get("/midwife/{id}", makeHTTPHandlerFunc(s.handleGetMidwifeByID))
	r.Post("/midwife", makeHTTPHandlerFunc(s.handleCreateMidwife))

	log.Printf("EvesTracker API is running on port: %s", s.listenAddr)

	http.ListenAndServe(s.listenAddr, r)
}

func (s *APIServer) handleGetMidwifeByID(w http.ResponseWriter, r *http.Request) *APIError {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}
	midwife, err := s.store.GetMidwifeByID(id)
	if err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}

	writeJSON(w, http.StatusOK, midwife)
	return nil
}

func (s *APIServer) handleCreateMidwife(w http.ResponseWriter, r *http.Request) *APIError {
	createMidwifeReq := new(CreateMidwifeRequest)
	if err := json.NewDecoder(r.Body).Decode(createMidwifeReq); err != nil {
		return &APIError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "bad request data",
		}
	}

	_, err := s.store.CreateMidwife(createMidwifeReq)
	if err != nil {
		return &APIError{
			Code:         http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
	}

	writeJSON(w, http.StatusOK, createMidwifeReq)
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
