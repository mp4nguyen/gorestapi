package personMdl

import "time"

type Person struct {
	PersonId       int       `json:"personId"`
	IsEnable       int       `json:"isEnable"`
	Title          string    `json:"title"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	Dob            time.Time `json:"dob"`
	Gender         string    `json:"gender"`
	Phone          string    `json:"phone"`
	Mobile         string    `json:"mobile"`
	Occupation     string    `json:"occupation"`
	Address        string    `json:"address"`
	SuburbDistrict string    `json:"suburbDistrict"`
	Ward           string    `json:"ward"`
	Postcode       string    `json:"postcode"`
	StateProvince  string    `json:"stateProvince"`
	Country        string    `json:"country"`
	IsPatient      int       `json:"isPatient"`
	IsDoctor       int       `json:"isDoctor"`
	CreatedBy      int       `json:"createdBy"`
	CreationDate   time.Time `json:"creationDate"`
	LastUpdatedBy  int       `json:"lastUpdatedBy"`
	LastUpdateDate time.Time `json:"lastUpdateDate"`
	Email          string    `json:"email"`
	SourceId       int       `json:"sourceId"`
	AvatarId       int       `json:"avatarId"`
	SignatureId    int       `json:"signatureId"`
}

type Persons []Person
