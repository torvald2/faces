package services

import (
	"atbmarket.comfaceapp/models"
)

type faceRecognizer interface {
	GetUserIDByFace([]byte) (int, error)
	GetNewFaceDescriptor([]byte) ([]float32, error)
	GetShopId() int
}

type profileStore interface {
	GetProfileById(int) (models.Profile, error)
	CreateProfile(name string, image []byte, descriptor []float32, shop int) (profileId int, err error)
}

type imageGetter interface {
	GetImage(profileId int) (data []byte, err error)
}

type jornalRecorder interface {
	NewJornalRecord(oper models.JornalOperation) error
}
type logger interface {
	LogBadRequest(request models.BadRequest) error
}
