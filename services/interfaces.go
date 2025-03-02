package services

import (
	"io"
	"time"

	"github.com/torvald2/faces/models"
)

type ProfileStore interface {
	GetProfileById(int) (models.Profile, error)
	GetShopProfiles(shopId int) (profiles []models.Profile, err error)
	CreateProfile(name string, image []byte, descriptor []float32, shop int) (profileId int, err error)
	LogBadRequest(request models.BadRequest) error
}

type ImageGetter interface {
	GetImage(profileId int) (data []byte, err error)
	GetBadRequestImage(recordId int) (data []byte, err error)
}

type JornalRecorder interface {
	NewJornalRecord(oper models.JornalOperation) error
}
type Logger interface {
	LogBadRequest(request models.BadRequest) error
}

type FaceRecognizer interface {
	GetUserIDByFace(image []byte, requestId string) (userId int, err error)
}
type JornalGetter interface {
	GetJornalRecords(start, end time.Time) (data []models.JornalOperationDB, err error)
}

type BadRequestsGetter interface {
	GetBadRequests(dateFrom, dateTo time.Time) (data []models.BadRequest, err error)
}

type ReportSender = func(attatch io.Reader, attachName string, emails string) error
type SheetCreator = func(operations []models.JornalOperationDB) (io.Reader, error)
