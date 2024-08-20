package main

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
