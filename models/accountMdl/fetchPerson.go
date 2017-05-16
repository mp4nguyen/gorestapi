package accountMdl

import (
	"fmt"
	"strconv"

	"bitbucket.org/restapi/models/personMdl"
)

func (m *Accounts) FetchPerson() (err error) {

	foreignKeys := map[string]string{}
	whereCondition := "person_id in ("
	for _, row := range *m {

		_, ok := foreignKeys[strconv.Itoa(row.PersonId)]
		if !ok {
			fmt.Println(" row.PersonId = ", row.PersonId)
			foreignKeys[string(row.PersonId)] = string(row.PersonId)
			whereCondition = whereCondition + strconv.Itoa(row.PersonId) + ","
		}
	}

	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"

	persons, err := personMdl.MapFind(whereCondition, "person_id")

	for _, row := range *m {

		person, ok := persons[strconv.Itoa(row.PersonId)]

		if ok {
			row.Person = person

		}
	}

	//output, _ := json.Marshal(*m)
	//fmt.Println("\n\n\n Results = ", string(output))
	return err
}
