package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiServer struct {
	Address string
	Store   Storage
}

type ApiError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
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

func NewAPIServer(address string, store Storage) *ApiServer {
	return &ApiServer{
		Address: address,
		Store:   store,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account/{id}", makeHTTPFunc(s.handleAccount))
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
		return s.handleModifyAccount(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccount(w, r)
	}

	return fmt.Errorf("invalid Method: %v", r.Method)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	userId := vars["id"]
	account := NewAccount("Nitish", "Kumar")
	account.ID, _ = strconv.Atoi(userId)
	WriteJSON(w, http.StatusOK, account)
	return nil
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	accRequest := CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&accRequest); err != nil {
		return err
	}

	account := NewAccount(accRequest.FirstName, accRequest.LastName)
	if err := s.Store.CreateAccount(account); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusOK, account)
}

func (s *ApiServer) handleModifyAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	
}

func (s *ApiServer) handleTransferAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}
