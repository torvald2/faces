package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"

	log "atbmarket.comfaceapp/app_logger"
	"atbmarket.comfaceapp/services"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func GetImageHandler(ig services.ImageGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rid := getRequestID(r.Context())
		params := mux.Vars(r)
		id := params["id"]
		numId, err := strconv.Atoi(id)
		if err != nil {
			responseWithError(err, w)
			log.Logger.Error("Profile ID error",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.Error(err))
			return
		}
		image, err := ig.GetImage(numId)
		if err != nil {
			responseWithError(err, w)
			log.Logger.Error("Get Face error",
				zap.String("Method", r.Method),
				zap.String("URL", r.RequestURI),
				zap.String("RequestID", rid),
				zap.Error(err))
		} else {
			buf := bytes.NewBuffer(image)
			w.Header().Set("Content-Type", "image/jpeg")
			w.Header().Set("Content-Length", fmt.Sprintf("%v", len(image)))
			_, err = io.Copy(w, buf)
			if err != nil {
				log.Logger.Error("Image sending error",
					zap.String("Method", r.Method),
					zap.String("URL", r.RequestURI),
					zap.String("RequestID", rid),
					zap.Error(err))
			} else {
				log.Logger.Debug("Response ok",
					zap.String("Method", r.Method),
					zap.String("URL", r.RequestURI),
					zap.String("RequestID", rid))
			}

		}

	})
}
