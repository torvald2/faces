package handlers

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"github.com/google/uuid"
	log "github.com/torvald2/faces/app_logger"
)

type ReqIdKey string

const RequestIDString ReqIdKey = "RequestID"

func getReqIdMidelware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := uuid.New()
		ctx = context.WithValue(ctx, RequestIDString, id.String())
		r = r.WithContext(ctx)
		log.Logger.Debug("Getting new Request",
			zap.String("Request ID", id.String()),
			zap.String("Method", r.Method),
			zap.String("URL", r.RequestURI),
		)
		next.ServeHTTP(w, r)
	})
}
