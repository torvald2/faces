package services

import (
	"time"

	"github.com/torvald2/faces/models"
)

func GetBadRequests(start, end time.Time, rg BadRequestsGetter) (data []models.BadRequest, err error) {
	data, err = rg.GetBadRequests(start, end)
	return
}

func GetBadRequestImage(ig ImageGetter, profileId int) (image []byte, err error) {
	image, err = ig.GetBadRequestImage(profileId)
	return
}
