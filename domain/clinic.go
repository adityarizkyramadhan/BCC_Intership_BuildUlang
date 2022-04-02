package domain

import "gorm.io/gorm"

type Clinic struct {
	*gorm.Model
	NameClinic     string
	UsernameClinic string
	PasswordClinic string
	SpreadSheet    string
	Contact        string
	Address        string
	AtasNama       string
	NoRekening     string
}

type ClinicRepository interface {
	Create(*Clinic) (*Clinic, error)
	Delete(*Clinic) (*Clinic, error)
	Update(*Clinic) (*Clinic, error)
	FindAll() ([]*Clinic, error)
	FindById(id string) (*Clinic, error)
	FindByCity(city string) ([]*Clinic, error)
}
