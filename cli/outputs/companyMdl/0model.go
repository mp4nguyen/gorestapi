package companyMdl

import (
	"time"

	"bitbucket.org/restapi/models/clinicMdl"
)

type Company struct {
	CompanyId       int                `json:"companyId"`
	CompanyName     string             `json:"companyName"`
	IsEnable        int                `json:"isEnable"`
	Address         string             `json:"address"`
	SuburbDistrict  string             `json:"suburbDistrict"`
	Ward            string             `json:"ward"`
	Postcode        string             `json:"postcode"`
	StateProvince   string             `json:"stateProvince"`
	Country         string             `json:"country"`
	Description     string             `json:"description"`
	Policy          string             `json:"policy"`
	ConditionToBook string             `json:"conditionToBook"`
	LogoPath        string             `json:"logoPath"`
	CreatedBy       int                `json:"createdBy"`
	CreationDate    time.Time          `json:"creationDate"`
	LastUpdatedBy   int                `json:"lastUpdatedBy"`
	LastUpdateDate  time.Time          `json:"lastUpdateDate"`
	Clinic          []clinicMdl.Clinic `json:"clinics"`
}

type Companys []*Company
