package rostersByDateVMdl

import (
	"time"

	"bitbucket.org/restapi/models/clinicMdl"
)

type RostersByDateV struct {
	CompanyId     int              `json:"companyId" mysql:"company_id"`
	WorkingSiteId int              `json:"workingSiteId" mysql:"working_site_id"`
	BookingTypeId int              `json:"bookingTypeId" mysql:"booking_type_id"`
	RosterDate    time.Time        `json:"rosterDate" mysql:"roster_date"`
	Clinic        clinicMdl.Clinic `json:"clinics"`
}

type RostersByDateVs []*RostersByDateV
