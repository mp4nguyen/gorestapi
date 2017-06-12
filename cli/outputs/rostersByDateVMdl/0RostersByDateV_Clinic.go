/*Please copy the property below to main model file: 0model.go to extend the relationship*/
package rostersByDateVMdl

type RostersByDateV struct{
	Clinic clinicMdl.Clinic `json:"clinic"`
	}

