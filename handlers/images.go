package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"atbmarket.comfaceapp/services"
	"github.com/gorilla/mux"
)

func GetImageHandler(ig services.ImageGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]
		numId, err := strconv.Atoi(id)
		if err != nil {
			respDesc := fmt.Sprintf("Не верный формат ID профиля. %v", err)
			responseWithError(respDesc, w, http.StatusBadRequest)
			return
		}
		image, err := ig.GetImage(numId)
		if err != nil {
			respDesc := fmt.Sprintf("Проблема при получении лица профиля %v", err)
			responseWithError(respDesc, w, http.StatusInternalServerError)
		} else {
			buf := bytes.NewBuffer(image)
			w.Header().Set("Content-Type", "image/jpeg")
			w.Header().Set("Content-Length", fmt.Sprintf("%v", len(image)))
			_, err = io.Copy(w, buf)
			if err != nil {
				fmt.Print("Error") //Log
			}

		}

	})
}
