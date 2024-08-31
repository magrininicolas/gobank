package api

import (
	"net/http"

	"github.com/magrininicolas/gobank/models"
)

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	acc := models.NewAccount("Nicolas", "Pereira")
	return WriteJSON(w, http.StatusOK, acc)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, JSONResponse{"message": "change da world, my final message good bye"})
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
