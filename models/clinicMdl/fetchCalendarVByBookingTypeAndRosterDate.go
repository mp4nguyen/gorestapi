package clinicMdl

import (
	"strconv"
	"time"

	"bitbucket.org/restapi/models/calendarVMdl"
)

func (m *Clinics) FetchCalendarVByBookingTypeAndDate(bookingTypeId int, rosterDate time.Time) (err error) {
	whereCondition := " booking_type_id = " + strconv.Itoa(bookingTypeId) + " and roster_date = '" + rosterDate.Format("2006-01-02") + "' and clinic_id in ("
	for _, row := range *m {
		whereCondition = whereCondition + strconv.Itoa(row.ClinicId) + ","
	}
	whereCondition = whereCondition[0:len(whereCondition)-1] + ")"
	tempMapData, err := calendarVMdl.MapFind("ClinicId", whereCondition, "clinic_id")
	for _, row := range *m {
		tempData, ok := tempMapData[strconv.Itoa(row.ClinicId)]
		if ok {
			row.Slots = tempData
		}
	}
	return err
}

func (m *Clinic) FetchCalendarVByBookingTypeAndDate(bookingTypeId int, rosterDate time.Time) (err error) {
	whereCondition := " booking_type_id = " + strconv.Itoa(bookingTypeId) + " and roster_date = '" + rosterDate.Format("2006-01-02") + "' and clinic_id = " + strconv.Itoa(m.ClinicId)
	tempMapData, err := calendarVMdl.MapFind("ClinicId", whereCondition, "clinic_id")

	tempData, ok := tempMapData[strconv.Itoa(m.ClinicId)]
	if ok {
		m.Slots = tempData
	}

	return err
}
