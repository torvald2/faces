package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"atbmarket.comfaceapp/services"
	"github.com/gorilla/mux"
)

func GetRecognizeFaceHandler(recognizers map[int]services.FaceRecognizer, ps services.ProfileStore, log services.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		numId, err := strconv.Atoi(id)
		if err != nil {
			respDesc := fmt.Sprintf("Shop ID convertation failed. %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}

		recognizer, ok := recognizers[numId]
		if !ok {
			respDesc := fmt.Sprintf("Shop with ID %v not found", numId)
			responseWithError(respDesc, w, http.StatusNotFound)
			return
		}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respDesc := fmt.Sprintf("Failed to read body data  %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}

		profileId, err := recognizer.GetUserIDByFace(bodyBytes)
		if err != nil {
			respDesc := fmt.Sprintf("Recognize Error  %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}

		profile, err := ps.GetProfileById(profileId)
		if err != nil {
			respDesc := fmt.Sprintf("Profile request error  %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
			return
		}
		responseOk(w, profile)

	})

}
