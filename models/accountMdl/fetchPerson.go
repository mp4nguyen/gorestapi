package accountMdl

import (
	"strconv"

	"bitbucket.org/restapi/models/personMdl"
)

func (m *Accounts) FetchPerson() (err error) {
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
	tempMapData, err := personMdl.MapFind(whereCondition, "person_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.PersonId)]
		if ok {
			row.Person = tempData
		}
	}
	return err
}

// package accountMdl
//
// import (
// 	"strconv"
//
// 	"bitbucket.org/restapi/models/personMdl"
// )
//
// func (m *Accounts) FetchPerson() (err error) {
// 	foreignKeys := map[string]string{}
// 	whereCondition := "person_id in ("
// 	for _, row := range *m {
// 		_, ok := foreignKeys[strconv.Itoa(row.PersonId)]
// 		if !ok {
// 			foreignKeys[strconv.Itoa(row.PersonId)] = strconv.Itoa(row.PersonId)
// 			whereCondition = whereCondition + strconv.Itoa(row.PersonId) + ","
// 		}
// 	}
// 	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
// 	tempMapData, err := personMdl.MapFind(whereCondition, "person_id")
// 	for _, row := range *m {
// 		tempData, ok := tempMapData[strconv.Itoa(row.PersonId)]
// 		if ok {
// 			row.Person = tempData
// 		}
// 	}
// 	return err
// }
