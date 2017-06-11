package onlineBookingCtrl

import (
	"net/http"

	"bitbucket.org/restapi/models/bookingTypeMdl"
	"bitbucket.org/restapi/utils"
)

func GetBookingTypes(w http.ResponseWriter, r *http.Request) {

	bts, err := bookingTypeMdl.Find(" 1 = 1 ", "1 ")
	utils.APIResponse(w, err, bts)

}
