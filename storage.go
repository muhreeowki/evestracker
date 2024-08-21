package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// TODO: CRUD functions for Midwife and Patients
type Storage interface {
	// Midwife Functions
	GetMidwifeByID(int) (*Midwife, error)
	CreateMidwife(*CreateMidwifeRequest) (*Midwife, error)
	DeleteMidwifeByID(int) error
	UpdateMidwifeByID(int) (*Midwife, error)
	GetMidwifeMothers(int) ([]*Mother, error)
	// Mother Functions
	GetMotherByID(int) *Mother
	CreateMother(Midwife) (*Mother, error)
	DeleteMotherByID(int) error
	UpdateMotherByID(int) (*Mother, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := os.Getenv("POSTGRES_CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to DB: %s", err.Error())
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	if err := s.CreateMotherTable(); err != nil {
		return err
	}
	if err := s.CreateMidwifeTable(); err != nil {
		return err
	}
	if err := s.AddPGCrypto(); err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) AddPGCrypto() error {
	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto`
	res, err := s.db.Exec(query)
	if err != nil {
		return err
	}

	fmt.Println("AddPGCrypto res: ", res)

	return nil
}

// Table Functions
func (s *PostgresStore) CreateMotherTable() error {
	query := `CREATE TABLE IF NOT EXISTS mother (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    firstname TEXT,
    lastname TEXT,
    birth_date TIMESTAMP,
    email TEXT,
    phone TEXT,
    address TEXT,
    partner_name TEXT,
    image_url TEXT,
    lmp TIMESTAMP,
    conception_date TIMESTAMP,
    sono_date TIMESTAMP,
    crl FLOAT,
    crl_date TIMESTAMP,
    edd TIMESTAMP,
    rh_factor TEXT,
    delivered BOOLEAN
  )`

	res, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("could not create mother table: %s", err)
	}

	fmt.Println("Mother table res: ", res)

	return nil
}

// Table Functions
func (s *PostgresStore) CreateMidwifeTable() error {
	query := `CREATE TABLE IF NOT EXISTS midwife (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    pass TEXT NOT NULL,
    image_url TEXT,
    mothers INTEGER[]
  )`

	res, err := s.db.Exec(query)
	if err != nil {
		return fmt.Errorf("could not create midwife table: %s", err)
	}

	fmt.Println("Midwife table res: ", res)

	return nil
}

// Midwife Functions
func (s *PostgresStore) GetMidwifeByID(id int) (*Midwife, error) {
	query := `SELECT * FROM midwife WHERE id == $1 LIMIT 1`
	row := s.db.QueryRow(query, id)

	midwife := new(Midwife)
	if err := row.Scan(midwife); err != nil {
		return nil, fmt.Errorf("no account found with id %d", id)
	}

	return midwife, nil
}

func (s *PostgresStore) CreateMidwife(midwife *CreateMidwifeRequest) (*Midwife, error) {
	query := `
  INSERT INTO midwife 
  (firstname, lastname, email, pass)
  VALUES ($1, $2, $3, crypt($4, gen_salt('bf')))
  `
	_, err := s.db.Query(query, midwife.FirstName, midwife.LastName, midwife.Email, midwife.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to create midwife: %s", err.Error())
	}
	return nil, nil
}

func (s *PostgresStore) DeleteMidwifeByID(int) error              { return nil }
func (s *PostgresStore) UpdateMidwifeByID(int) (*Midwife, error)  { return nil, nil }
func (s *PostgresStore) GetMidwifeMothers(int) ([]*Mother, error) { return nil, nil }

// Mother Functions
func (s *PostgresStore) GetMotherByID(int) *Mother             { return nil }
func (s *PostgresStore) CreateMother(Midwife) (*Mother, error) { return nil, nil }
func (s *PostgresStore) DeleteMotherByID(int) error            { return nil }
func (s *PostgresStore) UpdateMotherByID(int) (*Mother, error) { return nil, nil }
