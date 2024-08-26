package main

import (
	"database/sql"
	"time"
)

type CreateMidwifeRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" validate:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type Midwife struct {
	ID        uint           `db:"id"          json:"id"`
	FirstName string         `db:"firstname"   json:"firstName" binding:"required"`
	LastName  string         `db:"lastname"    json:"lastName"  binding:"required"`
	Email     string         `db:"email"       json:"email"     binding:"required" validate:"email"`
	Password  string         `db:"pass"        json:"password"  binding:"required"`
	ImageURL  sql.NullString `db:"image_url"   json:"imageURL"`
	// Mothers   []sql.NullInt64 `db:"mothers"     json:"mothers"`
	CreatedAt *time.Time   `db:"created_at"  json:"createdAt"`
	UpdatedAt sql.NullTime `db:"updated_at"  json:"updatedAt"`
}

type Mother struct {
	ID             uint            `db:"id"           json:"id"`
	CreatedAt      time.Time       `db:"created_at"   json:"createdAt"`
	UpdatedAt      sql.NullTime    `db:"updated_at"   json:"updatedAt"`
	DeletedAt      sql.NullTime    `db:"deleted_at"   json:"deletedAt,omitempty"`
	FirstName      string          `db:"firstname"    json:"firstName"    binding:"required"`
	LastName       sql.NullString  `db:"lastname"     json:"lastName,omitempty"`
	BirthDate      sql.NullTime    `db:"birth_date"   json:"birthDate,omitempty"`
	Email          sql.NullString  `db:"email"        json:"email,omitempty"`
	Phone          sql.NullString  `db:"phone"        json:"phone,omitempty"`
	Address        sql.NullString  `db:"address"      json:"address,omitempty"`
	PartnerName    sql.NullString  `db:"partner_name" json:"partnerName,omitempty"`
	ImageURL       sql.NullString  `db:"image_url"    json:"imageURL,omitempty"`
	LMP            sql.NullTime    `db:"lmp"          json:"lmp,omitempty"`
	ConceptionDate sql.NullTime    `db:"conception_date" json:"conceptionDate,omitempty"`
	SonoDate       sql.NullTime    `db:"sono_date"    json:"sonoDate,omitempty"`
	CRL            sql.NullFloat64 `db:"crl"         json:"crl,omitempty"`
	CRLDate        sql.NullTime    `db:"crl_date"     json:"crlDate,omitempty"`
	EDD            sql.NullTime    `db:"edd"          json:"edd,omitempty"`
	RhFactor       sql.NullString  `db:"rh_factor"    json:"rhFactor,omitempty"`
	Delivered      sql.NullBool    `db:"delivered"    json:"delivered,omitempty"`
	DeliveryDate   sql.NullTime    `db:"delivery_date" json:"deliveryDate,omitempty"`
	MidwifeID      sql.NullInt64   `db:"midwife_id"   json:"midwifeID,omitempty"`
}

type CreateMotherRequest struct {
	FirstName      string     `json:"firstName" binding:"required"`
	LastName       string     `json:"lastName"`
	BirthDate      *time.Time `json:"birthDate"`
	Email          string     `json:"email"`
	Phone          string     `json:"phone"`
	Address        string     `json:"address"`
	PartnerName    string     `json:"partnerName"`
	ImageURL       string     `json:"imageURL"`
	LMP            *time.Time `json:"lmp"`
	ConceptionDate *time.Time `json:"conceptionDate"`
	SonoDate       *time.Time `json:"sonoDate"`
	CRL            *float64   `json:"crl"`
	CRLDate        *time.Time `json:"crlDate"`
	EDD            *time.Time `json:"edd"`
	RhFactor       string     `json:"rhFactor"`
	MidwifeID      *uint32    `json:"midwifeID"`
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
