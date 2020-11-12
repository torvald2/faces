package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"atbmarket.comfaceapp/models"
)

func responseWithError(errorDesc string, w http.ResponseWriter, status int) {
	var respData models.HttpResponse
	respData.Error = errorDesc
	respData.Status = "ERROR"
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respData)
}

func responseOk(w http.ResponseWriter, data interface{}) {
	var respData models.HttpResponse
	respData.Status = "OK"
	respData.Data = data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respData)
}

func getRequestID(ctx context.Context) string {

	reqID := ctx.Value(RequestIDString)

	if ret, ok := reqID.(string); ok {
		return ret
	}

	return ""
}
