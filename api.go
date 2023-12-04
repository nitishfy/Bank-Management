package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiServer struct {
	Address string
}

type ApiError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

// makeHTTPFunc converts the api Function into handler func so that the[ handler can accept it
func makeHTTPFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func NewAPIServer(address string) *ApiServer {
	return &ApiServer{
		Address: address,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPFunc(s.handleAccount))
	log.Println("Server listening on port", s.Address)
	http.ListenAndServe(s.Address, router)
}

func (s *ApiServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "PUT" {
		return s.handleModifyAcount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("invalid Method: %v", r.Method)
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("Nitish", "Kumar")
	WriteJSON(w, http.StatusOK, account)
	return nil
}

func (s *ApiServer) handleModifyAcount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
