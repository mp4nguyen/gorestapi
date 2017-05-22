package accountMdl

import "log"
import "bitbucket.org/restapi/db"

func (m *Account)FetchPersonForAccount()(err error){
	whereCondition := "person_id = strconv.Itoa(m.PersonId)"
	tempMapData, err := personMdl.MapFind("PersonId",whereCondition, "person_id")
		tempData, ok := tempMapData[strconv.Itoa(row.PersonId)]
		if ok {
			if len(tempData) > 0 {
			m.Person = tempData[0]
			}
		}
	return err
}
func (m *Accounts)FetchPersonForAccounts()(err error){
	foreignKeys := map[string]string{}
	whereCondition := "person_id in ("
	for _, row := range *m {
		_, ok := foreignKeys[strconv.Itoa(row.PersonId)]
		if !ok {
			foreignKeys[strconv.Itoa(row.PersonId)] = strconv.Itoa(row.PersonId)
			whereCondition = whereCondition + strconv.Itoa(row.PersonId) + ","
		}
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := personMdl.MapFind("PersonId",whereCondition, "person_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.PersonId)]
		if ok {
			if len(tempData) > 0 {
			row.Person = tempData[0]
			}
		}
	}
	return err
}
