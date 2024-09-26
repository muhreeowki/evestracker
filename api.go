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
	// Midwife endpoints
	r.Get("/midwife", makeHTTPHandlerFunc(s.handleGetMidwives)) // NOTE: This must be deleted or protected fully. Nobody should access this route.

	r.Post("/midwife", makeHTTPHandlerFunc(s.handleCreateMidwife))
	r.Get("/midwife/{id}", makeHTTPHandlerFunc(s.handleGetMidwifeByID))
	r.Get("/midwife/{id}/mothers", makeHTTPHandlerFunc(s.handleGetMidwifeMothers))
	r.Delete("/midwife/{id}", makeHTTPHandlerFunc(s.handleDeleteMidwifeByID))

	// Mother endpoints
	r.Get("/mother", makeHTTPHandlerFunc(s.handleGetMothers)) // NOTE: This must be deleted or protected fully. Nobody should access this route.

	r.Post("/mother", makeHTTPHandlerFunc(s.handleCreateMother))
	r.Get("/mother/{id}", makeHTTPHandlerFunc(s.handleGetMotherByID))
	r.Delete("/mother/{id}", makeHTTPHandlerFunc(s.handleDeleteMotherByID))

	log.Printf("EvesTracker API is running on port: %s", s.listenAddr)
	http.ListenAndServe(s.listenAddr, r)
}

//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ MIDWIFE HANDLERS ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

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

func (s *APIServer) handleGetMidwifeMothers(w http.ResponseWriter, r *http.Request) *APIError {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}

	mothers, err := s.store.GetMidwifeMothers(id)
	if err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}

	return writeJSON(w, http.StatusOK, mothers)
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

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ MOTHER HANDLERS ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

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
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}

	if err := s.store.DeleteMotherByID(id); err != nil {
		return &APIError{
			ErrorMessage: "invalid id",
			Code:         http.StatusBadRequest,
		}
	}

	return writeJSON(w, http.StatusOK, fmt.Sprintf("successfully deleted mother of id %d", id))
}
