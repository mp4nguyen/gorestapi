package accountMdl

import "time"

type Account struct{
	Password string `json:"password" mysql:"password"`
	Email string `json:"email" mysql:"email"`
	UserType string `json:"userType" mysql:"user_type"`
	IsEnable int `json:"isEnable" mysql:"isEnable"`
	CreatedBy int `json:"createdBy" mysql:"created_by"`
	CreationDate time.Time `json:"creationDate" mysql:"creation_date"`
	LastUpdatedBy int `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
	PersonId int `json:"personId" mysql:"person_id"`
	DoctorId int `json:"doctorId" mysql:"doctor_id"`
	PatientId int `json:"patientId" mysql:"patient_id"`
	CompanyId int `json:"companyId" mysql:"company_id"`
	EmailVerified int `json:"emailVerified" mysql:"emailVerified"`
	Realm string `json:"realm" mysql:"realm"`
	Credentials string `json:"credentials" mysql:"credentials"`
	Challenges string `json:"challenges" mysql:"challenges"`
	VerificationToken string `json:"verificationToken" mysql:"verificationToken"`
	Status string `json:"status" mysql:"status"`
	Created time.Time `json:"created" mysql:"created"`
	Lastupdated time.Time `json:"lastupdated" mysql:"lastupdated"`
	Id int `json:"id" mysql:"id"`
	Username string `json:"username" mysql:"username"`
	}

type Accounts []*Account