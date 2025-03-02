package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"unicode/utf8"

	"go.uber.org/zap"

	"github.com/gorilla/mux"
	log "github.com/torvald2/faces/app_logger"
	"github.com/torvald2/faces/services"
)

func GetRecognizeFaceHandler(ps services.ProfileStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		rid := getRequestID(r.Context())
		numId, err := strconv.Atoi(id)
		if err != nil {
			responseWithError(err, w)
			log.Logger.Error("Bad shopID param",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid))
			return
		}

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responseWithError(err, w)
			log.Logger.Error("Bad shopID param. Is not integer",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid))
			return
		}
		profile, err := services.RecognizeFace(ps, bodyBytes, rid, numId)
		profile.ImageUrl = fmt.Sprintf("/images/%v", profile.Id)
		if err != nil {
			log.Logger.Warn("Face recognition error",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.Error(err))
			responseWithError(err, w)
		} else {
			responseOk(w, profile)
			log.Logger.Debug("Response ok",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid))
		}

	})

}

func GetNewFaceHandler(ps services.ProfileStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		rid := getRequestID(r.Context())
		id := params["id"]
		userName := r.FormValue("name")
		numId, err := strconv.Atoi(id)
		if err != nil {
			responseWithError(err, w)
			log.Logger.Error("Bad Shop Num",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid))
			return
		}
		if utf8.RuneCountInString(userName) < 1 {
			responseWithError(err, w)
			log.Logger.Error("Bad user FIO",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid))
			return
		}

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			responseWithError(err, w)
			log.Logger.Error("Failed to read body data",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.Error(err))
			return
		}
		profile, err := services.CreateNewProfile(ps, bodyBytes, userName, numId, rid)
		if err != nil {
			responseWithError(err, w)
			log.Logger.Warn("Face recognition problem",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.Error(err))
		} else {
			responseOk(w, profile)
			log.Logger.Debug("Face recognized",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid))
		}

	})

}
