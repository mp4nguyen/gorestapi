package accountMdl

type Login struct {
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PatientProfile struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Email           string `json:"email"`
	FatherPersonId  int    `json:"fatherPersonId"`
	PersonId        int    `json:"personId"`
	Title           string `json:"title"`
	FirstName       string `json:"firstName"`
	LastName        string `json:"lastName"`
	Dob             string `json:"dob"`
	Gender          string `json:"gender"`
	Occupation      string `json:"occupation"`
	Mobile          string `json:"phone"`
	Address         string `json:"address"`
	SuburbDistrict  string `json:"suburbDistrict"`
	StateProvince   string `json:"stateProvince"`
	Postcode        string `json:"postcode"`
	Country         string `json:"country"`
	GPFirstName     string `json:"gpFirstName"`
	GPLastName      string `json:"gpLastName"`
	Clinic          string `json:"clinic"`
	GPContactNumber string `json:"gpContactNumber"`
	MedicareNo      string `json:"medicareNo"`
	MedicareRef     string `json:"medicareRef"`
	MedicareExpired string `json:"medicareExpired"`
}

type signup struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type baseinfo struct {
	FatherPersonId int    `json:"fatherPersonId"`
	PersonId       int    `json:"personId"`
	Title          string `json:"title"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Dob            string `json:"dob"`
	Gender         string `json:"gender"`
	Occupation     string `json:"occupation"`
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
