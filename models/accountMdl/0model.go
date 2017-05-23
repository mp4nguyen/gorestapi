package accountMdl

import (
	"time"

	"bitbucket.org/restapi/models/personMdl"
)

type Login struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Account struct {
	Password          string           `json:"-"`
	Email             string           `json:"email"`
	UserType          string           `json:"userType"`
	IsEnable          int              `json:"isEnable"`
	CreatedBy         int              `json:"createdBy"`
	CreationDate      time.Time        `json:"creationDate"`
	LastUpdatedBy     int              `json:"lastUpdatedBy"`
	LastUpdateDate    time.Time        `json:"lastUpdateDate"`
	PersonId          int              `json:"personId"`
	DoctorId          int              `json:"doctorId"`
	PatientId         int              `json:"patientId"`
	CompanyId         int              `json:"companyId"`
	EmailVerified     int              `json:"emailVerified"`
	Realm             string           `json:"realm"`
	Credentials       string           `json:"credentials"`
	Challenges        string           `json:"challenges"`
	VerificationToken string           `json:"verificationToken"`
	Status            string           `json:"status"`
	Created           time.Time        `json:"created"`
	Lastupdated       time.Time        `json:"lastupdated"`
	Id                int              `json:"id"`
	Username          string           `json:"username"`
	Person            personMdl.Person `json:"person"`
	AccessToken       string           `json:"accessToken"`
}

type Accounts []*Account

type CheckAvailableRes struct {
	IsAvailable bool   `json:"isAvailable"`
	Reason      string `json:"reason"`
}

type signup struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type baseinfo struct {
	Title      string `json:"title"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Dob        string `json:"dob"`
	Gender     string `json:"gender"`
	Occupation string `json:"occupation"`
}

type contact struct {
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Suburb   string `json:"suburb"`
	State    string `json:"state"`
	Postcode string `json:"postcode"`
	Country  string `json:"country"`
}

type gp struct {
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Clinic          string `json:"clinic"`
	ContactNumber   string `json:"contactNumber"`
	MedicareNo      string `json:"medicareNo"`
	MedicareRef     string `json:"medicareRef"`
	MedicareExpired string `json:"medicareExpired"`
}

type Member struct {
	Signup   signup   `json:"signup"`
	Baseinfo baseinfo `json:"baseinfo"`
	Contact  contact  `json:"contact"`
	GP       gp       `json:"gp"`
}
