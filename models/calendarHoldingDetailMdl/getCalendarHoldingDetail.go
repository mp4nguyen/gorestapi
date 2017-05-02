package calendarHoldingDetailMdl

import (
	"log"
	"time"

	"bitbucket.org/restapi/db"
)

func GetCalendarHoldingDetail() (calendars CalendarHoldingDetails, err error) {

	start := time.Now()

	rows, err := db.GetDB().Query("select id,holding_id,cal_id,Creation_date,Last_update_date,socket_id,candidate_temp_id from calendar_holding_details ")

	if err != nil {
		log.Println("users.go: All() err = ", err)
	}

	Response := CalendarHoldingDetails{}

	for rows.Next() {

		calendar := CalendarHoldingDetail{}
		rows.Scan(&calendar.Id, &calendar.HoldingId, &calendar.CalId, &calendar.CreationDate, &calendar.LastUpdateDate, &calendar.SocketId, &calendar.CandidateTempId)
		//calendar.CalendarFromTimeInTime,err := time.Parse(layout, calendar.CalendarFromTime)
		Response.CalendarHoldingDetails = append(Response.CalendarHoldingDetails, calendar)
	}

	log.Printf("calendarHoldingDetailMdl: sql with normal way duration = %s", time.Since(start))

	return Response, err
}
