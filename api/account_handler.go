package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/magrininicolas/gobank/models"
)

func (s *APIServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return err
	}
	acc, err := s.dbx.GetAccountById(id)
	if err != nil {
		return err
	}
	WriteJSON(w, http.StatusOK, models.FromAccToGetResponse(acc))
	return err
}

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accs, err := s.dbx.GetAccounts()
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, &APIError{
			StatusCode: http.StatusBadRequest,
			Status:     "BAD REQUEST",
			Error:      err.Error(),
		})
		return err
	}

	accsResp := []*models.GetAccountResponse{}
	for _, v := range accs {
		accsResp = append(accsResp, models.FromAccToGetResponse(v))
	}
	WriteJSON(w, http.StatusOK, accsResp)
	return nil
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
