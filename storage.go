package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func (pg *PostgresStore) Init() error {
	return pg.createAccountTable()
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "postgres://postgres:123456@localhost:5432/postgres?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := dbConn.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{db: dbConn}, nil
}

func (pg *PostgresStore) CreateAccount(acc *Account) error {
	return nil
}

func (pg *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (pg *PostgresStore) UpdateAccount(acc *Account) error {
	return nil
}

func (pg *PostgresStore) GetAccountByID(int) (*Account, error) {
	return nil, nil
}

func (pg *PostgresStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key,
		fist_name varchar(50),
		last_name varchar(50),
		number VARCHAR(20),
		balance NUMBER(10, 2),
		created_at timestamp
	)`

	_, err := pg.db.Exec(query)
	return err
}
