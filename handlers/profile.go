package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	log "atbmarket.comfaceapp/app_logger"
	"atbmarket.comfaceapp/services"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func GetProfileHandler(ps services.ProfileStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		rid := getRequestID(r.Context())
		id := params["id"]
		numId, err := strconv.Atoi(id)
		if err != nil {
			respDesc := fmt.Sprintf("Не верный формат ID магазина. %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			log.Logger.Error("Bad shop ID",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.Error(err))
			return
		}
		profiles, err := services.GetProfiles(ps, numId)
		if err != nil {
			respDesc := fmt.Sprintf("Проблема при получении списка профилей %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
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
