package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"atbmarket.comfaceapp/models"

	"atbmarket.comfaceapp/services"
)

func GetWorkJornalHandler(wj services.JornalRecorder, log services.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var record models.JornalOperation

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&record)
		if err != nil {
			respDesc := fmt.Sprintf("Не верный формат сообщений %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
			return
		}

		err = services.RegisterJornalOperation(wj, log, record)
		if err != nil {
			respDesc := fmt.Sprintf("Проблема при регистрации записи журнала %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
			return
		} else {
			responseOk(w, nil)
		}

	})
}
