package personMdl

import "log"
import "bitbucket.org/restapi/db"

func (m *Person)FetchPatientRelationshipV()(err error){
%!(EXTRA string=Person)	whereCondition := "father_person_id = " + strconv.Itoa(m.PersonId)
	tempMapData, err := patientRelationshipVMdl.MapFind("FatherPersonId",whereCondition, "father_person_id")
		tempData, ok := tempMapData[strconv.Itoa(m.PersonId)]
		if ok {
			m.Relationships = tempData
		}
	return err
}
func (m *Persons)FetchPatientRelationshipV()(err error){
%!(EXTRA string=Person)	whereCondition := "father_person_id in ("
	for _, row := range *m {
			whereCondition = whereCondition + strconv.Itoa(row.PersonId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := patientRelationshipVMdl.MapFind("FatherPersonId",whereCondition, "father_person_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.PersonId)]
		if ok {
			row.Relationships = tempData
		}
	}
	return err
}
