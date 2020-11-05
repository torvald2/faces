package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"atbmarket.comfaceapp/services"
	"github.com/gorilla/mux"
)

func GetProfileHandler(ps services.ProfileStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		numId, err := strconv.Atoi(id)
		if err != nil {
			respDesc := fmt.Sprintf("Не верный формат ID магазина. %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}
		profiles, err := services.GetProfiles(ps, numId)
		if err != nil {
			respDesc := fmt.Sprintf("Проблема при получении списка профилей %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
		} else {
			responseOk(w, profiles)

		}

	})
}
