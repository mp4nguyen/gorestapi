package patientMdl

import "time"

type Patient struct{
	PatientId int `json:"patientId"`
	MedicalCompanyId int `json:"medicalCompanyId"`
	UserId int `json:"userId"`
	PersonId int `json:"personId"`
	IsEnable int `json:"isEnable"`
	CreatedBy int `json:"createdBy"`
	CreationDate time.Time `json:"creationDate"`
	LastUpdatedBy int `json:"lastUpdatedBy"`
	LastUpdateDate time.Time `json:"lastUpdateDate"`
	SourceId int `json:"sourceId"`
	FatherPatientId int `json:"fatherPatientId"`
	}

type Patients []*Patient