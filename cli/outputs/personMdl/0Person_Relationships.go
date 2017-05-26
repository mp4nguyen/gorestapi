/*Please copy the property below to main model file: 0model.go to extend the relationship*/
package personMdl

type Person struct{
	Relationships []patientRelationshipVMdl.PatientRelationshipV `json:"relationshipss"`
	}

