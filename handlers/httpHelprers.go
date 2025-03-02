package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/torvald2/faces/models"
	"github.com/torvald2/faces/services"
)

func responseWithError(err error, w http.ResponseWriter) {
	var statusCode int
	switch err.(type) {
	case services.NoFaceError:
		statusCode = 412
	case services.MultipleFaces:
		statusCode = 412
	case services.UserNotFound:
		statusCode = 404
	case services.NoProfilesForShop:
		statusCode = 404
	default:
		statusCode = 500
	}
	var respData models.HttpResponse
	respData.Error = fmt.Sprintf("%v", err)
	respData.Status = "ERROR"
	w.WriteHeader(statusCode)
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
