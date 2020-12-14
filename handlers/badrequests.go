package handlers

import (
	"net/http"
	"strconv"
	"time"

	"atbmarket.comfaceapp/services"
)

func GetBadRequestHandler(br services.BadRequestsGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startDateTimestamp, err := strconv.ParseInt(r.FormValue("start"), 10, 64)
		if err != nil {
			responseWithError(err, w)
			return
		}
		endDateTimestamp, err := strconv.ParseInt(r.FormValue("end"), 10, 64)
		if err != nil {
			responseWithError(err, w)
			return
		}

		endDate := time.Unix(endDateTimestamp, 0)
		startDate := time.Unix(startDateTimestamp, 0)
		data, err := services.GetBadRequests(startDate, endDate, br)
		if err != nil {
			responseWithError(err, w)
		} else {
			responseOk(w, data)
		}

	})
}
