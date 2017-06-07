package accountMdl

import (
	"errors"
	"time"

	"bitbucket.org/restapi/db"
	"bitbucket.org/restapi/logger"
	"bitbucket.org/restapi/models/accessTokenMdl"
	"bitbucket.org/restapi/utils"
	"github.com/go-sql-driver/mysql"
)

func (m Login) CheckAccount() (isMatch bool, account LoginRes, err error) {
	log := logger.Log
	loginRes := LoginRes{}
	start := time.Now()
	sqlString := "select password,email,user_type,isEnable,created_by,creation_date,last_updated_by,last_update_date,person_id,doctor_id,patient_id,company_id,emailVerified,realm,credentials,challenges,verificationToken,status,created,lastupdated,id,username from ocs.accounts where username=?"

	rs := db.GetDB().QueryRow(sqlString, m.Username)

	log.Infof(
		"sql duration = %s",
		time.Since(start),
	)
	row := Account{}
	tempCreationDate := mysql.NullTime{}
	tempLastUpdateDate := mysql.NullTime{}
	tempCreated := mysql.NullTime{}
	tempLastupdated := mysql.NullTime{}
	rs.Scan(&row.Password, &row.Email, &row.UserType, &row.IsEnable, &row.CreatedBy, &tempCreationDate, &row.LastUpdatedBy, &tempLastUpdateDate, &row.PersonId, &row.DoctorId, &row.PatientId, &row.CompanyId, &row.EmailVerified, &row.Realm, &row.Credentials, &row.Challenges, &row.VerificationToken, &row.Status, &tempCreated, &tempLastupdated, &row.Id, &row.Username)
	row.CreationDate = tempCreationDate.Time
	row.LastUpdateDate = tempLastUpdateDate.Time
	row.Created = tempCreated.Time
	row.Lastupdated = tempLastupdated.Time
	log.Infof(
		"sql duration2 = %s",
		time.Since(start),
	)
	if row.Password == "" {
		return false, loginRes, errors.New("Username does not exist")
	}

	isMatchAcc := utils.CheckPasswordHash(m.Password, row.Password)

	log.Infof(
		"sql duration3 = %s",
		time.Since(start),
	)
	//err = bcrypt.CompareHashAndPassword([]byte(row.Password), []byte(m.Password))

	if isMatchAcc == false {
		return false, loginRes, errors.New("Wrong password")
	} else {
		row.FetchPerson()
		row.Person.FetchPatientRelationshipV()
		//row.Person.FetchPatientAppointment()
		//row.Person.Relationships.FetchPatientAppointment()
		at, err := accessTokenMdl.Create(row.Id)
		utils.ErrorHandler("Accesstoken generated ", err, nil)

		row.AccessToken = at
		log.Infof(" at = %s", at)
		/////prepare object to return to client
		patientAccRes := PatientAccountRes{
			PersonId:       row.PersonId,
			PatientId:      row.PatientId,
			Username:       row.Username,
			Title:          row.Person.Title,
			FirstName:      row.Person.FirstName,
			LastName:       row.Person.LastName,
			Dob:            row.Person.Dob,
			Gender:         row.Person.Gender,
			Address:        row.Person.Address,
			SuburbDistrict: row.Person.SuburbDistrict,
			Ward:           row.Person.Ward,
			Postcode:       row.Person.Postcode,
			StateProvince:  row.Person.StateProvince,
			Country:        row.Person.Country,
			Profile:        row.Person,
			AccessToken:    row.AccessToken,
		}

		loginRes.IsLogin = true
		loginRes.AccessToken = at
		loginRes.Account = patientAccRes

		return true, loginRes, err
	}

}
