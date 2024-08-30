package models

import (
	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	AccNumber uuid.UUID `json:"accNumber"`
	Balance   float64   `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		AccNumber: uuid.New(),
	}
}
