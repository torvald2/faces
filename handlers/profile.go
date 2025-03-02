package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	log "github.com/torvald2/faces/app_logger"
	"github.com/torvald2/faces/services"
	"go.uber.org/zap"
)

func GetProfileHandler(ps services.ProfileStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		rid := getRequestID(r.Context())
		id := params["id"]
		numId, err := strconv.Atoi(id)
		if err != nil {
			responseWithError(err, w)
			log.Logger.Error("Bad shop ID",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.Error(err))
			return
		}
		profiles, err := services.GetProfiles(ps, numId)
		if err != nil {
			responseWithError(err, w)
			log.Logger.Error("Get profiles error",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.Error(err))
		} else {
			responseOk(w, profiles)
			log.Logger.Debug("Get profiles error",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid))

		}

	})
}
