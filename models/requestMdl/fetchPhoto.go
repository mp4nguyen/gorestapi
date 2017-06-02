package requestMdl

import (
	"strconv"

	"bitbucket.org/restapi/models/photoMdl"
)

func (m *Request) FetchPhoto() (err error) {
	whereCondition := "request_id = " + strconv.Itoa(m.RequestId)
	tempMapData, err := photoMdl.MapFind("RequestId", whereCondition, "request_id")
	tempData, ok := tempMapData[strconv.Itoa(m.RequestId)]
	if ok {
		m.Photos = tempData
	}
	return err
}
func (m *Requests) FetchPhoto() (err error) {
	whereCondition := "request_id in ("
	for _, row := range *m {
		whereCondition = whereCondition + strconv.Itoa(row.RequestId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := photoMdl.MapFind("RequestId", whereCondition, "request_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.RequestId)]
		if ok {
			row.Photos = tempData
		}
	}
	return err
}
