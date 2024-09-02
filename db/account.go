package db

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
	row, err := s.dbx.Queryx("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		return scanIntoAccount(row)
	}
	return nil, fmt.Errorf("ACCOUNT WITH ID %v NOT FOUND", id)
}

func (s *PostgresDB) GetAccounts() ([]*models.Account, error) {
	rows, err := s.dbx.Queryx("SELECT * FROM account")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accs := []*models.Account{}
	for rows.Next() {
		acc, err := scanIntoAccount(rows)
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

func scanIntoAccount(rows *sqlx.Rows) (*models.Account, error) {
	acc := &models.Account{}
	err := rows.Scan(&acc.ID, &acc.FirstName, &acc.LastName, &acc.AccNumber, &acc.Balance, &acc.CreatedAt)
	if err != nil {
		return nil, err
	}
	return acc, nil
}
