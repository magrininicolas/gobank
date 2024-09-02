package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreateAccountResponse struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type GetAccountResponse struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type Account struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	AccNumber uuid.UUID `json:"accNumber"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		AccNumber: uuid.New(),
		CreatedAt: time.Now().UTC(),
	}
}

func FromAccToGetResponse(acc *Account) *GetAccountResponse {
	return &GetAccountResponse{
		FirstName: acc.FirstName,
		LastName:  acc.LastName,
		Balance:   acc.Balance,
		CreatedAt: acc.CreatedAt,
	}
}
