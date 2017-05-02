package calendarCtrl

import "bitbucket.org/restapi/models/calendarMdl"

type CalendarController struct{}

type GetCalendarParams struct {
	Id        int    `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	MaxPeriod int    `json:"maxPeriod"`
}

type returnCal struct {
	Dates []*returnDate `json:"calendars"`
}

type returnDate struct {
	Date         string        `json:"date"`
	FormatedDate string        `json:"formatedDate"`
	Doctors      []*calDoctor  `json:"doctors"`
	Slots        []*returnSlot `json:"slots"`
}

type calDoctor struct {
	DoctorID     int                    `json:"doctorID"`
	DoctorName   string                 `json:"doctorName"`
	Slots        []calendarMdl.Calendar `json:"-"`
	SlotSegments []*slotSegment         `json:"slotSegments"`
}

type slotSegment struct {
	Slots          []*calendarMdl.Calendar `json:"-"`
	FirstSlotBlock *slotBlock              `json:"firstSlot"`
	LastSlotBlock  *slotBlock              `json:"LastSlot"`
}

type slotBlock struct {
	FirstSlot calendarMdl.Calendar   `json:"slot"`
	Slots     []calendarMdl.Calendar `json:"followingSlots"`
}

type returnSlot struct {
	CalId          int                    `json:"calId"`
	CalendarDate   string                 `json:"calendarDate"`
	CalendarTime   string                 `json:"calendarTime"`
	DoctorId       int                    `json:"doctorId"`
	DoctorName     string                 `json:"doctorName"`
	FromTime       string                 `json:"fromTime"`
	ToTime         string                 `json:"toTime"`
	FromTimeInInt  int                    `json:"fromTimeInInt"`
	SiteId         int                    `json:"siteId"`
	FollowingSlots []calendarMdl.Calendar `json:"followingSlots"`
}

// Sort for returnSlot: ByTime implements sort.Interface for []returnSlot based on
// the CalendarTime field.
type ByCalendarTime []*returnSlot

func (a ByCalendarTime) Len() int           { return len(a) }
func (a ByCalendarTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCalendarTime) Less(i, j int) bool { return a[i].CalendarTime < a[j].CalendarTime }
