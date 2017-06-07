package personMdl

type CreatedMemberRes struct {
	IsSuccess bool        `json:"isSuccess"`
	Reason    string      `json:"reason"`
	PersonId  int         `json:"personId"`
	PatientId int         `json:"patientId"`
	Account   interface{} `json:"account"`
}
