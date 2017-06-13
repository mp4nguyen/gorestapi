package clinicMdl

import (
	"time"

	"bitbucket.org/restapi/models/calendarVMdl"
)

type Clinic struct {
	ClinicId       int                     `json:"clinicId"`
	ClinicName     string                  `json:"clinicName"`
	IsEnable       int                     `json:"isEnable"`
	CompanyId      int                     `json:"companyId"`
	IsBookable     int                     `json:"isBookable"`
	IsTelehealth   int                     `json:"isTelehealth"`
	IsCalendar     int                     `json:"isCalendar"`
	Description    string                  `json:"description"`
	Address        string                  `json:"address"`
	SuburbDistrict string                  `json:"suburbDistrict"`
	Ward           string                  `json:"ward"`
	Postcode       string                  `json:"postcode"`
	StateProvince  string                  `json:"stateProvince"`
	Country        string                  `json:"country"`
	CreatedBy      int                     `json:"createdBy"`
	CreationDate   time.Time               `json:"creationDate"`
	LastUpdatedBy  int                     `json:"lastUpdatedBy"`
	LastUpdateDate time.Time               `json:"lastUpdateDate"`
	Latitude       float64                 `json:"latitude"`
	Longitude      float64                 `json:"longitude"`
	IconBase64     string                  `json:"iconBase64" mysql:"icon_base64"`
	Slots          calendarVMdl.CalendarVs `json:"slots"`
}

type Clinics []*Clinic
