package personMdl

import "time"

type Person struct{
	PersonId int `json:"personId" mysql:"person_id"`
	IsEnable int `json:"isEnable" mysql:"isEnable"`
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
	IsPatient int `json:"isPatient" mysql:"isPatient"`
	IsDoctor int `json:"isDoctor" mysql:"isDoctor"`
	CreatedBy int `json:"createdBy" mysql:"created_by"`
	CreationDate time.Time `json:"creationDate" mysql:"creation_date"`
	LastUpdatedBy int `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
	Email string `json:"email" mysql:"email"`
	SourceId int `json:"sourceId" mysql:"source_id"`
	AvatarId int `json:"avatarId" mysql:"avatar_id"`
	SignatureId int `json:"signatureId" mysql:"signature_id"`
	GPFirstName string `json:"gPFirstName" mysql:"GP_First_name"`
	GPLastName string `json:"gPLastName" mysql:"GP_Last_name"`
	ClinicName string `json:"clinicName" mysql:"Clinic_Name"`
	GPContact string `json:"gPContact" mysql:"GP_Contact"`
	MedicareNo string `json:"medicareNo" mysql:"Medicare_No"`
	MedicareRef string `json:"medicareRef" mysql:"Medicare_ref"`
	MedicareExpired time.Time `json:"medicareExpired" mysql:"Medicare_Expired"`
	}

type Persons []*Person