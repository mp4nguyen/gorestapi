package patientVMdl

import "time"

type PatientV struct{
	PatientId int `json:"patientId" mysql:"patient_id"`
	MedicalCompanyId int `json:"medicalCompanyId" mysql:"medical_company_id"`
	PersonId int `json:"personId" mysql:"person_id"`
	IsEnable int `json:"isEnable" mysql:"isEnable"`
	CreatedBy int `json:"createdBy" mysql:"created_by"`
	CreationDate time.Time `json:"creationDate" mysql:"creation_date"`
	LastUpdatedBy int `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
	Title string `json:"title" mysql:"title"`
	FirstName string `json:"firstName" mysql:"first_name"`
	LastName string `json:"lastName" mysql:"last_name"`
	Dob time.Time `json:"dob" mysql:"dob"`
	Gender string `json:"gender" mysql:"gender"`
	Phone string `json:"phone" mysql:"phone"`
	Mobile string `json:"mobile" mysql:"mobile"`
	Occupation string `json:"occupation" mysql:"occupation"`
	Address string `json:"address" mysql:"address"`
	SuburbDistrict string `json:"suburbDistrict" mysql:"suburb_district"`
	Ward string `json:"ward" mysql:"ward"`
	Postcode string `json:"postcode" mysql:"postcode"`
	StateProvince string `json:"stateProvince" mysql:"state_province"`
	Country string `json:"country" mysql:"country"`
	Email string `json:"email" mysql:"email"`
	AvatarId int `json:"avatarId" mysql:"avatar_id"`
	AvatarUrl string `json:"avatarUrl" mysql:"avatar_url"`
	SignatureId int `json:"signatureId" mysql:"signature_id"`
	SignatureUrl string `json:"signatureUrl" mysql:"signature_url"`
	}

type PatientVs []*PatientV