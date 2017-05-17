/*Please copy the property below to main model file: 0model.go to extend the relationship*/
package companyMdl

type Company struct{
	Clinics []clinicMdl.Clinic `json:"clinicss"`
	}

