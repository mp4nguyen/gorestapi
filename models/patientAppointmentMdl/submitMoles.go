package patientAppointmentMdl

import (
	"database/sql"
	"encoding/json"
	"time"

	"bitbucket.org/restapi/db"
	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/photoMdl"
	"bitbucket.org/restapi/models/requestMdl"
	"bitbucket.org/restapi/utils"
)

func (m MoleRequest) SubmitMoles() error {
	log := logger.Log
	err := db.Transaction(func(tx *sql.Tx) error {
		creatingAppt := PatientAppointment{}
		creatingAppt.ApptDate = time.Now().UTC()
		creatingAppt.ApptStatus = "Confirmed"
		creatingAppt.ApptType = "MOLEPATROL"
		creatingAppt.PatientPersonId = m.PersonId
		creatingAppt.PatientId = m.PatientId
		creatingAppt.RequireDate = time.Now().UTC()
		noOfAppt, apptId, apptErr := creatingAppt.Create(tx)
		utils.LogError("Failed to create person", apptErr)
		if apptErr != nil {
			return apptErr
		}
		log.Infof("noOfAppt=%s, apptId=%s, apptErr=%s", noOfAppt, apptId, apptErr)

		for _, lesion := range m.Lesions {
			output, err := json.Marshal(lesion)
			utils.LogError("marshal lesion Json ", err)
			creatingReq := requestMdl.Request{}
			creatingReq.ApptId = int(apptId)
			creatingReq.PersonId = m.PersonId
			creatingReq.PatientId = m.PatientId
			creatingReq.Type = "MOLEPATROL"
			creatingReq.Data = string(output)
			noOfReq, reqId, reqErr := creatingReq.Create(tx)
			utils.LogError("Failed to create person", reqErr)
			if reqErr != nil {
				return reqErr
			}
			log.Infof("noOfReq=%s, reqId=%s, reqErr=%s", noOfReq, reqId, reqErr)

			creatingPhotos := photoMdl.Photos{}
			for _, photo := range lesion.Photos {
				creatingPhoto := photoMdl.Photo{
					RequestId: int(reqId),
					ApptId:    int(apptId),
					PatientId: m.PatientId,
					PersonId:  m.PersonId,
					Type:      "MOLEPATROL",
					Uri:       photo,
				}
				creatingPhotos = append(creatingPhotos, &creatingPhoto)
			}
			if len(lesion.Lesion) > 0 {
				creatingPhoto := photoMdl.Photo{
					RequestId: int(reqId),
					ApptId:    int(apptId),
					PatientId: m.PatientId,
					PersonId:  m.PersonId,
					Type:      "MOLEPATROL",
					Uri:       lesion.Lesion,
				}
				creatingPhotos = append(creatingPhotos, &creatingPhoto)
			}

			noOfPhoto, _, photoErr := creatingPhotos.Create(tx)
			utils.LogError("Failed to create photos", photoErr)
			if photoErr != nil {
				return photoErr
			}
			log.Infof("noOfPhoto=%s, photoErr=%s", noOfPhoto, photoErr)

		}
		return nil
	})
	return err
}
