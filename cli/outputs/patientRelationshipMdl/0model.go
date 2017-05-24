package patientRelationshipMdl

import "time"

type PatientRelationship struct{
	RelationshipId int `json:"relationshipId"`
	PatientId int `json:"patientId"`
	PersonId int `json:"personId"`
	RelationshipType string `json:"relationshipType"`
	IsEnable int `json:"isEnable"`
	CreatedBy int `json:"createdBy"`
	CreationDate time.Time `json:"creationDate"`
	LastUpdatedBy int `json:"lastUpdatedBy"`
	LastUpdateDate time.Time `json:"lastUpdateDate"`
	FatherPersonId int `json:"fatherPersonId"`
	}

type PatientRelationships []*PatientRelationship