package main

import (
	"time"
)

type Midwife struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" validate:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	ImageURL  string `json:"profileImage"`
	Mothers   []Mother
}

type Mother struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time
	// Patient's personal details
	FirstName   string    `json:"firstName"`
	LastName    string    `json:"lastName"`
	BirthDate   time.Time `json:"birthDate"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	PartnerName string    `json:"partnerName"`
	ImageURL    string    `json:"imageURL"`
	// Patients' medical details
	LMP            time.Time `json:"lmp"`            // Last Menstrual Period
	ConceptionDate time.Time `json:"conceptionDate"` // Date of conception
	SonoDate       time.Time `json:"sonoDate"`       // Date of sonogram
	CRL            float64   `json:"crl"`            // Crown Rump Length
	CRLDate        time.Time `json:"crlDate"`        // Date of CRL
	EDD            time.Time `json:"edd"`            // Estimated Due Date
	RhFactor       string    `json:"rhFactor"`       // Rh Factor
	// Delivery details
	Delivered    bool      `json:"delivered"`    // Has the patient delivered
	DeliveryDate time.Time `json:"deliveryDate"` // Date of delivery
	// Midwife details
	MidwifeID uint32 `json:"midwifeID"`
}
