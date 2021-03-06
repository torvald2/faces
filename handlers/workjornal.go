package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

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
			responseWithError(err, w)
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
			responseWithError(err, w)
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

func GetSendWorkJornalHandler(jg services.JornalGetter, rs services.ReportSender, sh services.SheetCreator) http.Handler {
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
		emails := r.FormValue("emails")

		endDate := time.Unix(endDateTimestamp, 0)
		startDate := time.Unix(startDateTimestamp, 0)
		err = services.SendJornalByMail(startDate, endDate, emails, jg, rs, sh)
		if err != nil {
			responseWithError(err, w)
		} else {
			responseOk(w, "OK")
		}

	})
}
