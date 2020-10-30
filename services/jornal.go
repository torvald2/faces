package services

import (
	"atbmarket.comfaceapp/models"
)

func RegisterJornalOperation(j jornalRecorder, log logger, operation models.JornalOperation) error {
	if operation.UserId != operation.RecognizedUserID {
		go func() {
			log.LogBadRequest(models.BadRequest{
				RequestId:       operation.RequestId,
				ErrorType:       models.NotMe,
				UserId:          operation.UserId,
				RecognizedUsers: []int{operation.RecognizedUserID},
			})

		}()
	}
	err := j.NewJornalRecord(operation)
	return err
}
