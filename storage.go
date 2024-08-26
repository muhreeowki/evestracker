package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	_ "github.com/lib/pq"
)

// TODO: CRUD functions for Midwife and Patients
type Storage interface {
	// Midwife Functions
	CreateMidwife(*CreateMidwifeRequest) (*Midwife, error)
	GetMidwives() ([]*Midwife, error)
	GetMidwifeByID(int) (*Midwife, error)
	GetMidwifeMothers(int) ([]*Mother, error)
	UpdateMidwifeByID(int) (*Midwife, error)
	DeleteMidwifeByID(int) error
	// Mother Functions
	CreateMother(*CreateMotherRequest) (*Mother, error)
	GetMothers() ([]*Mother, error)
	GetMotherByID(int) *Mother
	UpdateMotherByID(int) (*Mother, error)
	DeleteMotherByID(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := os.Getenv("POSTGRES_CONNSTR")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("could not connect to DB: %s", err.Error())
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	if err := s.CreateMotherTable(); err != nil {
		log.Println(err.Error())
		return err
	}
	if err := s.CreateMidwifeTable(); err != nil {
		log.Println(err.Error())
		return err
	}
	if err := s.AddPGCrypto(); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s *PostgresStore) AddPGCrypto() error {
	query := `CREATE EXTENSION IF NOT EXISTS pgcrypto`
	res, err := s.db.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	fmt.Println("AddPGCrypto res: ", res)

	return nil
}

// Table Functions
func (s *PostgresStore) CreateMidwifeTable() error {
	query := `CREATE TABLE IF NOT EXISTS midwife (
    id SERIAL PRIMARY KEY,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    pass TEXT NOT NULL,
    image_url TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
  )`

	res, err := s.db.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("could not create midwife table: %s", err)
	}

	fmt.Println("Midwife table res: ", res)

	return nil
}

// GetMidwifeByID gets a midwife from the database with the matching id
func (s *PostgresStore) GetMidwifeByID(id int) (*Midwife, error) {
	query := `SELECT * FROM midwife WHERE id = $1 LIMIT 1`
	row := s.db.QueryRow(query, id)

	midwife := new(Midwife)
	midwifeFields := getFields(midwife)
	if err := row.Scan(midwifeFields...); err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("no account found with id %d", id)
	}

	return midwife, nil
}

func (s *PostgresStore) GetMidwives() ([]*Midwife, error) {
	query := `SELECT * FROM midwife`

	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("could not get midwives")
	}

	midwives := []*Midwife{}
	for rows.Next() {
		midwife := new(Midwife)
		midwifeFields := getFields(midwife)

		if err := rows.Scan(midwifeFields...); err != nil {
			log.Println(err.Error())
			return nil, fmt.Errorf("could not get midwives")
		}

		midwives = append(midwives, midwife)
	}

	return midwives, nil
}

func (s *PostgresStore) CreateMidwife(midwife *CreateMidwifeRequest) (*Midwife, error) {
	query := `
  INSERT INTO midwife 
  (firstname, lastname, email, pass, created_at)
  VALUES ($1, $2, $3, crypt($4, gen_salt('bf')), $5)
  `
	_, err := s.db.Query(query, midwife.FirstName, midwife.LastName, midwife.Email, midwife.Password, time.Now())
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("failed to create midwife: %s", err.Error())
	}
	return nil, nil
}

func (s *PostgresStore) DeleteMidwifeByID(id int) error {
	query := `DELETE FROM midwife WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("no account found with id %d", id)
	}
	return nil
}

func (s *PostgresStore) UpdateMidwifeByID(int) (*Midwife, error) { return nil, nil }
func (s *PostgresStore) GetMidwifeMothers(id int) (midwives []*Mother, err error) {
	_, err = s.GetMidwifeByID(id) // Midwife
	if err != nil {
		return nil, err
	}
	return midwives, nil
}

// Mother Functions
func (s *PostgresStore) CreateMotherTable() error {
	query := `CREATE TABLE IF NOT EXISTS mother (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    firstname TEXT NOT NULL,
    lastname TEXT,
    birth_date TIMESTAMP WITH TIME ZONE,
    email TEXT,
    phone TEXT,
    address TEXT,
    partner_name TEXT,
    image_url TEXT,
    lmp TIMESTAMP WITH TIME ZONE,
    conception_date TIMESTAMP WITH TIME ZONE,
    sono_date TIMESTAMP WITH TIME ZONE,
    crl FLOAT,
    crl_date TIMESTAMP WITH TIME ZONE,
    edd TIMESTAMP WITH TIME ZONE,
    rh_factor TEXT,
    delivered BOOLEAN,
    delivery_date TIMESTAMP WITH TIME ZONE,
    midwife_id INTEGER
  )`

	res, err := s.db.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("could not create mother table: %s", err)
	}

	fmt.Println("Mother table res: ", res)

	return nil
}

func (s *PostgresStore) GetMothers() ([]*Mother, error) {
	query := `SELECT * FROM mother`

	rows, err := s.db.Query(query)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("could not get mothers")
	}

	mothers := []*Mother{}
	for rows.Next() {
		mother := new(Mother)
		motherFields := getFields(mother)

		if err := rows.Scan(motherFields...); err != nil {
			log.Println(err.Error())
			return nil, fmt.Errorf("could not get midwives")
		}

		mothers = append(mothers, mother)
	}

	return mothers, nil
}

func (s *PostgresStore) GetMotherByID(int) *Mother { return nil }

func (s *PostgresStore) CreateMother(mother *CreateMotherRequest) (*Mother, error) {
	query := `
  INSERT INTO mother (
    created_at, firstname, lastname, birth_date,
    email, phone, address, partner_name, image_url,
    lmp, conception_date, sono_date, crl, crl_date,
    edd, rh_factor, midwife_id
  ) VALUES (
    CURRENT_TIMESTAMP, $1, $2, $3, $4, $5, $6, $7,
    $8, $9, $10, $11, $12, $13, $14, $15, $16
  ) RETURNING *`
	_, err := s.db.Query(query,
		mother.FirstName,
		nullString(mother.LastName),
		nullTime(mother.BirthDate),
		nullString(mother.Email),
		nullString(mother.Phone),
		nullString(mother.Address),
		nullString(mother.PartnerName),
		nullString(mother.ImageURL),
		nullTime(mother.LMP),
		nullTime(mother.ConceptionDate),
		nullTime(mother.SonoDate),
		nullFloat64(mother.CRL),
		nullTime(mother.CRLDate),
		nullTime(mother.EDD),
		nullString(mother.RhFactor),
		nullInt32(mother.MidwifeID))
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("failed to create midwife: %s", err.Error())
	}
	return nil, nil
}

func (s *PostgresStore) DeleteMotherByID(int) error            { return nil }
func (s *PostgresStore) UpdateMotherByID(int) (*Mother, error) { return nil, nil }

func getFields(v any) []interface{} {
	s := reflect.ValueOf(v).Elem()
	numFields := s.NumField()
	fields := make([]interface{}, numFields)

	for i := 0; i < numFields; i++ {
		field := s.Field(i)
		fields[i] = field.Addr().Interface()
	}

	return fields
}

func nullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

func nullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func nullFloat64(f *float64) sql.NullFloat64 {
	if f == nil {
		return sql.NullFloat64{}
	}
	return sql.NullFloat64{Float64: *f, Valid: true}
}

func nullInt32(i *uint32) sql.NullInt32 {
	if i == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: int32(*i), Valid: true}
}
