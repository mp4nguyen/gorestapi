package patientAppointmentMdl

type MoleRequest struct {
	PersonId  int      `json:"personId"`
	PatientId int      `json:"patientId"`
	Lesions   []Lesion `json:"lesions"`
}

type Lesion struct {
	LesionId             int      `json:"lesionId"`
	Lesion               string   `json:"lesion"`
	IsFront              bool     `json:"isFront"`
	IsNew                bool     `json:"isNew"`
	IsGrowing            bool     `json:"isGrowing"`
	IsShapeOrChangeColor bool     `json:"isShapeOrChangeColor"`
	IsItchyOrBleeding    bool     `json:"isItchyOrBleeding"`
	IsTenderOrPainful    bool     `json:"isTenderOrPainful"`
	DoesItComeAndGo      bool     `json:"doesItComeAndGo"`
	Photos               []string `json:"photos"`
}
