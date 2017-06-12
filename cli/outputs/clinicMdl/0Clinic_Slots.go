/*Please copy the property below to main model file: 0model.go to extend the relationship*/
package clinicMdl

type Clinic struct{
	Slots calendarVMdl.CalendarVs `json:"slotss"`
	}

