package bookingTypeMdl

import "time"

type BookingType struct{
	BookingTypeId int `json:"bookingTypeId" mysql:"booking_type_id"`
	BookingTypeName string `json:"bookingTypeName" mysql:"booking_type_name"`
	IsEnable int `json:"isEnable" mysql:"isEnable"`
	CreatedBy int `json:"createdBy" mysql:"created_by"`
	CreationDate time.Time `json:"creationDate" mysql:"creation_date"`
	LastUpdatedBy int `json:"lastUpdatedBy" mysql:"last_updated_by"`
	LastUpdateDate time.Time `json:"lastUpdateDate" mysql:"last_update_date"`
	Icon text `json:"icon" mysql:"icon"`
	}

type BookingTypes []*BookingType