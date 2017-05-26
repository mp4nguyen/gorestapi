package accountMdl

import (
	"time"

	"bitbucket.org/restapi/models/personMdl"
)

type PatientAccountRes struct {
	PersonId       int              `json:"personId"`
	Username       string           `json:"username"`
	Title          string           `json:"title"`
	FirstName      string           `json:"firstName"`
	LastName       string           `json:"lastName"`
	Dob            time.Time        `json:"dob"`
	Gender         string           `json:"gender"`
	Address        string           `json:"address"`
	SuburbDistrict string           `json:"suburbDistrict"`
	Ward           string           `json:"ward"`
	Postcode       string           `json:"postcode"`
	StateProvince  string           `json:"stateProvince"`
	Country        string           `json:"country"`
	Profile        personMdl.Person `json:"profile"`
	AccessToken    string           `json:"accessToken"`
	// GPFirstName     string           `json:"gpFirstName"`
	// GPLastName      string           `json:"gpLastName"`
	// Clinic          string           `json:"clinic"`
	// GPContactNumber string           `json:"gpContactNumber"`
	// MedicareNo      string           `json:"medicareNo"`
	// MedicareRef     string           `json:"medicareRef"`
	// MedicareExpired time.Time        `json:"medicareExpired"`
}

type CheckAvailableRes struct {
	IsAvailable bool   `json:"isAvailable"`
	Reason      string `json:"reason"`
}

type LoginRes struct {
	IsLogin     bool              `json:"isLogin"`
	Reason      string            `json:"reason"`
	AccessToken string            `json:"accessToken"`
	Account     PatientAccountRes `json:"account"`
}
