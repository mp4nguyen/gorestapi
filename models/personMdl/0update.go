package personMdl

import "time"

type UpdatingMember struct {
	PersonId        int       `json:"personId" mysql:"person_id PRI"`
	Title           string    `json:"title" mysql:"title"`
	FirstName       string    `json:"firstName" mysql:"first_name"`
	LastName        string    `json:"lastName" mysql:"last_name"`
	Dob             time.Time `json:"dob" mysql:"dob"`
	Gender          string    `json:"gender" mysql:"gender"`
	Phone           string    `json:"phone" mysql:"phone"`
	Mobile          string    `json:"mobile" mysql:"mobile"`
	Occupation      string    `json:"occupation" mysql:"occupation"`
	Address         string    `json:"address" mysql:"address"`
	SuburbDistrict  string    `json:"suburbDistrict" mysql:"suburb_district"`
	Postcode        string    `json:"postcode" mysql:"postcode"`
	StateProvince   string    `json:"stateProvince" mysql:"state_province"`
	Country         string    `json:"country" mysql:"country"`
	LastUpdatedBy   int       `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate  time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
	Email           string    `json:"email" mysql:"email"`
	GPFirstName     string    `json:"gPFirstName" mysql:"GP_First_name"`
	GPLastName      string    `json:"gPLastName" mysql:"GP_Last_name"`
	ClinicName      string    `json:"clinicName" mysql:"Clinic_Name"`
	GPContact       string    `json:"gPContact" mysql:"GP_Contact"`
	MedicareNo      string    `json:"medicareNo" mysql:"Medicare_No"`
	MedicareRef     string    `json:"medicareRef" mysql:"Medicare_ref"`
	MedicareExpired time.Time `json:"medicareExpired" mysql:"Medicare_Expired"`
}

// PersonId        int    `json:"personId" mysql:"person_id"`
// Title           string `json:"title" mysql:"person_id"`
// FirstName       string `json:"firstName" mysql:"person_id"`
// LastName        string `json:"lastName" mysql:"person_id"`
// Dob             string `json:"dob" mysql:"person_id"`
// Gender          string `json:"gender" mysql:"person_id"`
// Occupation      string `json:"occupation" mysql:"person_id"`
// Phone           string `json:"phone" mysql:"person_id"`
// Address         string `json:"address" mysql:"person_id"`
// Suburb          string `json:"suburb" mysql:"person_id"`
// State           string `json:"state" mysql:"person_id"`
// Postcode        string `json:"postcode" mysql:"person_id"`
// Country         string `json:"country" mysql:"person_id"`
// GPFirstName     string `json:"firstName" mysql:"person_id"`
// GPLastName      string `json:"lastName" mysql:"person_id"`
// Clinic          string `json:"clinic" mysql:"person_id"`
// GPContactNumber string `json:"contactNumber" mysql:"person_id"`
// MedicareNo      string `json:"medicareNo" mysql:"person_id"`
// MedicareRef     string `json:"medicareRef" mysql:"person_id"`
// MedicareExpired string `json:"medicareExpired" mysql:"person_id"`
