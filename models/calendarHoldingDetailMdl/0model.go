package calendarHoldingDetailMdl

import "time"

type CalendarHoldingDetail struct {
	Id              int       `json:"id"`
	HoldingId       int       `json:"holdingId" sql:"roster_id"`
	CalId           int       `json:"calId" sql:"doctor_id"`
	CreationDate    time.Time `json:"creationDate" sql:"doctor_name"`
	LastUpdateDate  time.Time `json:"lastUpdateDate" sql:"calendar_from_time"`
	SocketId        int       `json:"socketId" sql:"calendar_to_time"`
	CandidateTempId int       `json:"candidateTempId" sql:"site_id"`
}

type CalendarHoldingDetails struct {
	CalendarHoldingDetails []CalendarHoldingDetail `json:"calendarHoldingDetails"`
}
