package patientRelationshipVMdl

import (
	"time"

	"bitbucket.org/restapi/models/patientAppointmentMdl"
)

type PatientRelationshipV struct {
	RelationshipId   int                                        `json:"relationshipId"`
	RelationshipType string                                     `json:"relationshipType"`
	PatientId        int                                        `json:"patientId"`
	PersonId         int                                        `json:"personId"`
	FatherPersonId   int                                        `json:"fatherPersonId"`
	IsEnable         int                                        `json:"isEnable"`
	CreatedBy        int                                        `json:"createdBy"`
	CreationDate     time.Time                                  `json:"creationDate"`
	LastUpdatedBy    int                                        `json:"lastUpdatedBy"`
	LastUpdateDate   time.Time                                  `json:"lastUpdateDate"`
	Title            string                                     `json:"title"`
	FirstName        string                                     `json:"firstName"`
	LastName         string                                     `json:"lastName"`
	Dob              time.Time                                  `json:"dob"`
	Gender           string                                     `json:"gender"`
	Phone            string                                     `json:"phone"`
	Mobile           string                                     `json:"mobile"`
	Occupation       string                                     `json:"occupation"`
	Address          string                                     `json:"address"`
	SuburbDistrict   string                                     `json:"suburbDistrict"`
	Ward             string                                     `json:"ward"`
	Postcode         string                                     `json:"postcode"`
	StateProvince    string                                     `json:"stateProvince"`
	Country          string                                     `json:"country"`
	Email            string                                     `json:"email"`
	AvatarId         int                                        `json:"avatarId"`
	AvatarUrl        string                                     `json:"avatarUrl"`
	SignatureId      int                                        `json:"signatureId"`
	SignatureUrl     string                                     `json:"signatureUrl"`
	GPFirstName      string                                     `json:"gPFirstName"`
	GPLastName       string                                     `json:"gPLastName"`
	ClinicName       string                                     `json:"clinicName"`
	GPContact        string                                     `json:"gPContact"`
	MedicareNo       string                                     `json:"medicareNo"`
	MedicareRef      string                                     `json:"medicareRef"`
	MedicareExpired  time.Time                                  `json:"medicareExpired"`
	Appointments     []patientAppointmentMdl.PatientAppointment `json:"appointments"`
}

type PatientRelationshipVs []*PatientRelationshipV
