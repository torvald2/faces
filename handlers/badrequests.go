package handlers

import (
	"net/http"
	"time"

	"github.com/torvald2/faces/services"
)

func GetBadRequestHandler(br services.BadRequestsGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startDate, err := time.Parse("2006-01-02", r.FormValue("start"))
		if err != nil {
			responseWithError(err, w)
			return
		}
		endDate, err := time.Parse("2006-01-02", r.FormValue("end"))
		if err != nil {
			responseWithError(err, w)
			return
		}

		data, err := services.GetBadRequests(startDate, endDate, br)
		if err != nil {
			responseWithError(err, w)
		} else {
			responseOk(w, data)
		}

	})
}
