package photoMdl

import "time"

type Photo struct {
	PhotoId        int       `json:"photoId" mysql:"photo_id"`
	RequestId      int       `json:"requestId" mysql:"request_id"`
	ApptId         int       `json:"apptId" mysql:"appt_id"`
	PatientId      int       `json:"patientId" mysql:"patient_id"`
	PersonId       int       `json:"personId" mysql:"person_id"`
	Type           string    `json:"type" mysql:"type"`
	Data           string    `json:"data" mysql:"data"`
	Uri            string    `json:"uri" mysql:"uri"`
	CreatedBy      int       `json:"createdBy" mysql:"created_by"`
	CreationDate   time.Time `json:"creationDate" mysql:"creation_date"`
	LastUpdatedBy  int       `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
}

type Photos []*Photo
