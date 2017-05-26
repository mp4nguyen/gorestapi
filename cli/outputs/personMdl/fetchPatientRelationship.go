package personMdl

import "log"
import "bitbucket.org/restapi/db"

func (m *Person)FetchPatientRelationship()(err error){
%!(EXTRA string=Person)	whereCondition := "father_person_id = " + strconv.Itoa(m.FatherPersonId)
	tempMapData, err := patientRelationshipMdl.MapFind(FatherPersonId,whereCondition, "father_person_id")
		tempData, ok := tempMapData[strconv.Itoa(m.FatherPersonId)]
		if ok {
			if len(tempData) > 0 {
			m.Relationships = tempData[0]
			}
		}
	return err
}
func (m *Persons)FetchPatientRelationship()(err error){
%!(EXTRA string=Person)	whereCondition := "father_person_id in ("
	for _, row := range *m {
			whereCondition = whereCondition + strconv.Itoa(row.PersonId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := patientRelationshipMdl.MapFind(FatherPersonId,whereCondition, "father_person_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.FatherPersonId)]
		if ok {
			row.Relationships = tempData
		}
	}
	return err
}
