package main

import (
	"encoding/json"
	"fmt"
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
	r.Get("/midwife", makeHTTPHandlerFunc(s.handleGetMidwives))
	r.Post("/midwife", makeHTTPHandlerFunc(s.handleCreateMidwife))
	r.Get("/midwife/{id}", makeHTTPHandlerFunc(s.handleGetMidwifeByID))
	r.Delete("/midwife/{id}", makeHTTPHandlerFunc(s.handleDeleteMidwifeByID))

	r.Post("/mother", makeHTTPHandlerFunc(s.handleCreateMother))
	r.Get("/mother", makeHTTPHandlerFunc(s.handleGetMothers))
	r.Get("/mother/{id}", makeHTTPHandlerFunc(s.handleGetMotherByID))

	log.Printf("EvesTracker API is running on port: %s", s.listenAddr)

	http.ListenAndServe(s.listenAddr, r)
}

func (s *APIServer) handleGetMidwives(w http.ResponseWriter, r *http.Request) *APIError {
	midwives, err := s.store.GetMidwives()
	if err != nil {
		return &APIError{
			ErrorMessage: err.Error(),
			Code:         http.StatusBadRequest,
		}
	}

	return writeJSON(w, http.StatusOK, midwives)
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

	return writeJSON(w, http.StatusOK, midwife)
}

func (s *APIServer) handleDeleteMidwifeByID(w http.ResponseWriter, r *http.Request) *APIError {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}

	if err := s.store.DeleteMidwifeByID(id); err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}

	return writeJSON(w, http.StatusOK, fmt.Sprintf("successfully deleted midwife of id %d", id))
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

	return writeJSON(w, http.StatusOK, createMidwifeReq)
}

func (s *APIServer) handleGetMothers(w http.ResponseWriter, r *http.Request) *APIError {
	mothers, err := s.store.GetMothers()
	if err != nil {
		return &APIError{
			ErrorMessage: err.Error(),
			Code:         http.StatusBadRequest,
		}
	}

	return writeJSON(w, http.StatusOK, mothers)
}

func (s *APIServer) handleCreateMother(w http.ResponseWriter, r *http.Request) *APIError {
	createMotherReq := new(CreateMotherRequest)
	if err := json.NewDecoder(r.Body).Decode(createMotherReq); err != nil {
		return &APIError{
			Code:         http.StatusBadRequest,
			ErrorMessage: "bad request data",
		}
	}

	_, err := s.store.CreateMother(createMotherReq)
	if err != nil {
		return &APIError{
			Code:         http.StatusBadRequest,
			ErrorMessage: err.Error(),
		}
	}

	return writeJSON(w, http.StatusOK, createMotherReq)
}

func (s *APIServer) handleGetMotherByID(w http.ResponseWriter, r *http.Request) *APIError {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}
	mother, err := s.store.GetMotherByID(id)
	if err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}

	return writeJSON(w, http.StatusOK, mother)
}

func (s *APIServer) handleDeleteMotherByID(w http.ResponseWriter, r *http.Request) *APIError {
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
