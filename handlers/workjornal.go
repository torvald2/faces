package handlers

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	applog "atbmarket.comfaceapp/app_logger"
	"go.uber.org/zap"

	"atbmarket.comfaceapp/services"
)

type DistanceRequest struct {
	Doc                string    `json:"doc"`
	Photo              string    `json:"photo"`
	ReturnVector       bool      `json:"return_vector"`
	ComparePhotoVector bool      `json:"compare_photo_vector"`
	Vector             []float64 `json:"vector"`
}

func DiscHandler(processor services.Descriptor) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var record DistanceRequest
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

		img1Bytes, err := base64.StdEncoding.DecodeString(record.Doc)
		if err != nil {

			responseWithError(err, w)

			return
		}
		if record.ReturnVector {
			dist1, err := processor.GetNewFaceDescriptor(img1Bytes)
			if err != nil {

				responseWithError(err, w)

				return
			}
			responseOk(w, map[string]interface{}{"data": dist1})
			return
		}
		if record.ComparePhotoVector {

			dist1, err := processor.GetNewFaceDescriptor(img1Bytes)
			if err != nil {

				responseWithError(err, w)

				return
			}

			dd, err := services.GetProfDistance(dist1, record.Vector)
			if err != nil {
				responseWithError(err, w)
				return
			}
			responseOk(w, map[string]interface{}{"dist": dd})
			return

		}

		img2Bytes, err := base64.StdEncoding.DecodeString(record.Photo)
		if err != nil {

			responseWithError(err, w)

			return
		}

		dist1, err := processor.GetNewFaceDescriptor(img1Bytes)
		if err != nil {

			responseWithError(err, w)

			return
		}

		dist2, err := processor.GetNewFaceDescriptor(img2Bytes)
		if err != nil {

			responseWithError(err, w)

			return
		}
		dd, err := services.GetProfDistance(dist1, dist2)

		if err != nil {

			responseWithError(err, w)

			return
		}

		responseOk(w, map[string]interface{}{"dist": dd})

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
