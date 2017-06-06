package patientAppointmentMdl

type MoleRequestRes struct {
	Lesions []LesionRes `json:"lesions"`
}

type LesionRes struct {
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
	Resource             []string `json:"resource"`
}
