package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	applog "atbmarket.comfaceapp/app_logger"
	"atbmarket.comfaceapp/models"
	"go.uber.org/zap"

	"atbmarket.comfaceapp/services"
)

func GetWorkJornalHandler(wj services.JornalRecorder, log services.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var record models.JornalOperation
		rid := getRequestID(r.Context())
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&record)
		if err != nil {
			raw_data := new(strings.Builder)
			io.Copy(raw_data, decoder.Buffered())
			respDesc := fmt.Sprintf("Не верный формат сообщений %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
			applog.Logger.Warn("Request decode error",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.String("data", raw_data.String()),
				zap.Error(err))
			return
		}

		err = services.RegisterJornalOperation(wj, log, record)
		if err != nil {
			request_data, _ := json.Marshal(record)
			respDesc := fmt.Sprintf("Проблема при регистрации записи журнала %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
			applog.Logger.Warn("Jornal registration error",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.String("Data", string(request_data)),
				zap.Error(err))
			return
		} else {
			responseOk(w, nil)
			applog.Logger.Debug("Response ok",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid))
		}

	})
}
