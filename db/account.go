package db

import (
	"github.com/google/uuid"
	"github.com/magrininicolas/gobank/models"
)

func (s *PostgresDB) createAccountTable() error {
	query := `
	CREATE TABLE if not exists account(
		id UUID primary key,
		first_name varchar(30) not null,
		last_name varchar(30) not null,
		acc_number UUID not null unique,
		balance numeric(15, 2) not null,
		created_at timestamp
	);
	`

	_, err := s.dbx.Exec(query)
	return err
}

func (s *PostgresDB) CreateAccount(acc *models.Account) error {
	query := `INSERT INTO account values($1, $2, $3, $4, $5, $6)`
	_, err := s.dbx.Exec(query, acc.ID, acc.FirstName, acc.LastName, acc.AccNumber, acc.Balance, acc.CreatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresDB) DeleteAccount(id uuid.UUID) error {
	return nil
}

func (s *PostgresDB) UpdateAccount(*models.Account) error {
	return nil
}

func (s *PostgresDB) GetAccountById(id uuid.UUID) (*models.Account, error) {
	return nil, nil
}

func (s *PostgresDB) GetAccounts() ([]models.Account, error) {
	rows, err := s.dbx.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accs := []models.Account{}
	for rows.Next() {
		acc := models.Account{}
		err := rows.Scan(&acc.ID, &acc.FirstName, &acc.LastName, &acc.AccNumber, &acc.Balance, &acc.CreatedAt)
		if err != nil {
			return nil, err
		}

		accs = append(accs, acc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accs, nil
}
