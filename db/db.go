package db

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/magrininicolas/gobank/models"
)

type Database interface {
	CreateAccount(*models.Account) error
	DeleteAccount(uuid.UUID) error
	UpdateAccount(*models.Account) error
	GetAccounts() ([]*models.Account, error)
	GetAccountById(uuid.UUID) (*models.Account, error)
}

type PostgresDB struct {
	dbx *sqlx.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	connStr := "user=admin dbname=gobank password=1234 sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{dbx: db}, nil
}

func (s *PostgresDB) Init() error {
	return s.createAccountTable()
}
