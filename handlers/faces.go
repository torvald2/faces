package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"atbmarket.comfaceapp/services"
	"github.com/ascarter/requestid"
	"github.com/gorilla/mux"
)

func GetRecognizeFaceHandler(ps services.ProfileStore, ra services.RecognizeAgregator) http.Handler {
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
		profile, err := services.RecognizeFace(ra, ps, bodyBytes, rid, numId)
		if err != nil {
			respDesc := fmt.Sprintf("Проблема при расспознании лица %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
		} else {
			responseOk(w, profile)
		}

	})

}

func GetNewFaceHandler(ps services.ProfileStore, ra services.DescriptorGetter) http.Handler {
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
		profile, err := services.RecognizeFace(ra, ps, bodyBytes, rid, numId)
		if err != nil {
			respDesc := fmt.Sprintf("Проблема при расспознании лица %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
		} else {
			responseOk(w, profile)
		}

	})

}
