package services

import (
	"atbmarket.comfaceapp/models"
)

type ProfileStore interface {
	GetProfileById(int) (models.Profile, error)
	GetShopProfiles(shopId int) (profiles []models.Profile, err error)
	CreateProfile(name string, image []byte, descriptor []float32, shop int) (profileId int, err error)
	LogBadRequest(request models.BadRequest) error
}

type ImageGetter interface {
	GetImage(profileId int) (data []byte, err error)
}

type JornalRecorder interface {
	NewJornalRecord(oper models.JornalOperation) error
}
type Logger interface {
	LogBadRequest(request models.BadRequest) error
}
