package services

import (
	"time"

	"github.com/torvald2/faces/models"
)

func RegisterJornalOperation(j JornalRecorder, log Logger, operation models.JornalOperation) error {
	if operation.UserId != operation.RecognizedUserID {
		go func() {
			log.LogBadRequest(models.BadRequest{
				RequestId:       operation.RequestId,
				ErrorType:       models.NotMe,
				UserId:          operation.UserId,
				RecognizedUsers: []int{operation.RecognizedUserID},
				RecognizeTime:   time.Now(),
			})

		}()
	}
	err := j.NewJornalRecord(operation)
	return err
}
