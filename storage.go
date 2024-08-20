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
	GetMidwifeByID(int) *Midwife
	CreateMidwife(Midwife) (*Midwife, error)
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

// Midwife Functions
func (s *PostgresStore) GetMidwifeByID(int) *Midwife              { return nil }
func (s *PostgresStore) CreateMidwife(Midwife) (*Midwife, error)  { return nil, nil }
func (s *PostgresStore) DeleteMidwifeByID(int) error              { return nil }
func (s *PostgresStore) UpdateMidwifeByID(int) (*Midwife, error)  { return nil, nil }
func (s *PostgresStore) GetMidwifeMothers(int) ([]*Mother, error) { return nil, nil }

// Mother Functions
func (s *PostgresStore) GetMotherByID(int) *Mother             { return nil }
func (s *PostgresStore) CreateMother(Midwife) (*Mother, error) { return nil, nil }
func (s *PostgresStore) DeleteMotherByID(int) error            { return nil }
func (s *PostgresStore) UpdateMotherByID(int) (*Mother, error) { return nil, nil }
