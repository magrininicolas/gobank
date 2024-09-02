package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/magrininicolas/gobank/models"
)

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	fmt.Println(id)
	return WriteJSON(w, http.StatusOK, &models.Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	req := models.CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	acc := models.NewAccount(req.FirstName, req.LastName)
	if err := s.dbx.CreateAccount(acc); err != nil {
		return err
	}

	return WriteJSON(w, http.StatusCreated, &models.CreateAccountResponse{
		FirstName: acc.FirstName,
		LastName:  acc.LastName,
		Balance:   acc.Balance,
		CreatedAt: acc.CreatedAt,
	})
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	WriteJSON(w, http.StatusOK, JSONResponse{"message": "transfer completed", "amount": 3000})
	return nil
}
