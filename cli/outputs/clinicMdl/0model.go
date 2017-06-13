package clinicMdl

import "time"

type Clinic struct {
	ClinicId       int       `json:"clinicId" mysql:"clinic_id"`
	ClinicName     string    `json:"clinicName" mysql:"clinic_name"`
	IsEnable       int       `json:"isEnable" mysql:"isEnable"`
	CompanyId      int       `json:"companyId" mysql:"company_id"`
	IsBookable     int       `json:"isBookable" mysql:"isBookable"`
	IsTelehealth   int       `json:"isTelehealth" mysql:"isTelehealth"`
	IsCalendar     int       `json:"isCalendar" mysql:"isCalendar"`
	Description    string    `json:"description" mysql:"description"`
	Address        string    `json:"address" mysql:"address"`
	SuburbDistrict string    `json:"suburbDistrict" mysql:"suburb_district"`
	Ward           string    `json:"ward" mysql:"ward"`
	Postcode       string    `json:"postcode" mysql:"postcode"`
	StateProvince  string    `json:"stateProvince" mysql:"state_province"`
	Country        string    `json:"country" mysql:"country"`
	CreatedBy      int       `json:"createdBy" mysql:"created_by"`
	CreationDate   time.Time `json:"creationDate" mysql:"creation_date"`
	LastUpdatedBy  int       `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
	Latitude       float     `json:"latitude" mysql:"latitude"`
	Longitude      float     `json:"longitude" mysql:"longitude"`
	IconBase64     string    `json:"iconBase64" mysql:"icon_base64"`
}

type Clinics []*Clinic
