package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"unicode/utf8"

	"atbmarket.comfaceapp/services"
	"github.com/ascarter/requestid"
	"github.com/gorilla/mux"
)

func GetRecognizeFaceHandler(ps services.ProfileStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		numId, err := strconv.Atoi(id)
		if err != nil {
			respDesc := fmt.Sprintf("Не верный формат номера магазина. %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respDesc := fmt.Sprintf("Failed to read body data  %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}
		rid, _ := requestid.FromContext(r.Context())
		profile, err := services.RecognizeFace(ps, bodyBytes, rid, numId)
		profile.ImageUrl = fmt.Sprintf("/images/%v", profile.Id)
		if err != nil {
			respDesc := fmt.Sprintf("Проблема при расспознании лица %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
		} else {
			responseOk(w, profile)
		}

	})

}

func GetNewFaceHandler(ps services.ProfileStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		userName := r.FormValue("name")
		numId, err := strconv.Atoi(id)
		if err != nil {
			respDesc := fmt.Sprintf("Не верный формат номера магазина. %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}
		if utf8.RuneCountInString(userName) < 1 {
			respDesc := fmt.Sprintf("Фио пользователя не должно быть пустым. %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}

		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respDesc := fmt.Sprintf("Failed to read body data  %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}
		rid, _ := requestid.FromContext(r.Context())
		profile, err := services.CreateNewProfile(ps, bodyBytes, userName, numId, rid)
		if err != nil {
			respDesc := fmt.Sprintf("Проблема при расспознании лица %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
		} else {
			responseOk(w, profile)
		}

	})

}
