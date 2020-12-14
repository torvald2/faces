package services

import (
	"time"

	"atbmarket.comfaceapp/models"
)

func GetBadRequests(start, end time.Time, rg BadRequestsGetter) (data []models.BadRequest, err error) {
	data, err = rg.GetBadRequests(start, end)
	return
}
