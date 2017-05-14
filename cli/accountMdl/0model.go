package accountMdl

import "time"

type Account struct {
	Password          string    `json:"password"`
	Email             string    `json:"email"`
	UserType          string    `json:"userType"`
	IsEnable          int       `json:"isEnable"`
	CreatedBy         int       `json:"createdBy"`
	CreationDate      time.Time `json:"creationDate"`
	LastUpdatedBy     int       `json:"lastUpdatedBy"`
	LastUpdateDate    time.Time `json:"lastUpdateDate"`
	PersonId          int       `json:"personId"`
	DoctorId          int       `json:"doctorId"`
	PatientId         int       `json:"patientId"`
	CompanyId         int       `json:"companyId"`
	EmailVerified     int       `json:"emailVerified"`
	Realm             string    `json:"realm"`
	Credentials       string    `json:"credentials"`
	Challenges        string    `json:"challenges"`
	VerificationToken string    `json:"verificationToken"`
	Status            string    `json:"status"`
	Created           time.Time `json:"created"`
	Lastupdated       time.Time `json:"lastupdated"`
	Id                int       `json:"id"`
	Username          string    `json:"username"`
}

type Accounts []Account
