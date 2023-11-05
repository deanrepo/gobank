package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	UpdateAccount(*Account) error
	DeleteAccount(int) error
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
	insertSql := "INSERT INTO account (first_name, last_name, number, balance, created_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := pg.db.Exec(insertSql, acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreateAt)

	return err
}

func (pg *PostgresStore) GetAccountByID(id int) (*Account, error) {
	var acc Account
	err := pg.db.QueryRow("SELECT id, first_name, last_name, number, balance, created_at FROM account WHERE id = $1", id).Scan(
		&acc.ID,
		&acc.FirstName,
		&acc.LastName,
		&acc.Number,
		&acc.Balance,
		&acc.CreateAt,
	)

	if err != nil && err == sql.ErrNoRows {
		log.Println("No record found")
		return nil, nil
	}

	return &acc, err
}

func (pg *PostgresStore) UpdateAccount(acc *Account) error {
	return nil
}

func (pg *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (pg *PostgresStore) createAccountTable() error {
	log.Println("Create table account...")
	query := `create table if not exists account (
		id serial primary key,
		first_name varchar(50),
		last_name varchar(50),
		number VARCHAR(20),
		balance NUMERIC(10, 2),
		created_at timestamp
	)`

	_, err := pg.db.Exec(query)
	return err
}
