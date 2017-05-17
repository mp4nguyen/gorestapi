package clinicMdl

import "time"

type Clinic struct {
	ClinicId       int       `json:"clinicId"`
	ClinicName     string    `json:"clinicName"`
	IsEnable       int       `json:"isEnable"`
	CompanyId      int       `json:"companyId"`
	IsBookable     int       `json:"isBookable"`
	IsTelehealth   int       `json:"isTelehealth"`
	IsCalendar     int       `json:"isCalendar"`
	Description    string    `json:"description"`
	Address        string    `json:"address"`
	SuburbDistrict string    `json:"suburbDistrict"`
	Ward           string    `json:"ward"`
	Postcode       string    `json:"postcode"`
	StateProvince  string    `json:"stateProvince"`
	Country        string    `json:"country"`
	CreatedBy      int       `json:"createdBy"`
	CreationDate   time.Time `json:"creationDate"`
	LastUpdatedBy  int       `json:"lastUpdatedBy"`
	LastUpdateDate time.Time `json:"lastUpdateDate"`
	Latitude       int64     `json:"latitude"`
	Longitude      int64     `json:"longitude"`
}

type Clinics []*Clinic
