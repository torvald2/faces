package services

import (
	"atbmarket.comfaceapp/adaptors"
	"atbmarket.comfaceapp/models"
)

func RecognizeFace(fr faceRecognizer, ps profileStore, log logger, image []byte, requestId string) (profile models.Profile, err error) {
	profileId, err := fr.GetUserIDByFace(image)

	if err != nil {
		go func() {
			br := models.BadRequest{
				CurrentFace: image,
				RequestId:   requestId,
				Shop:        fr.GetShopId(),
			}
			switch err.(type) {
			case *adaptors.NoFaceError:
				br.ErrorType = models.NoFace
			case *adaptors.MultipleFaces:
				br.ErrorType = models.MultipleRecognized
			case *adaptors.UserNotFound:
				br.ErrorType = models.UserNotFound

			}
			log.LogBadRequest(br)
		}()
		return
	}
	profile, err = ps.GetProfileById(profileId)
	return
}

func CreateNewProfile(fr faceRecognizer, ps profileStore, log logger, image []byte, name string, shop int, requestId string) (profileId int, err error) {
	descriptor, err := fr.GetNewFaceDescriptor(image)
	if err != nil {
		go func() {
			br := models.BadRequest{
				CurrentFace: image,
				RequestId:   requestId,
				Shop:        fr.GetShopId(),
			}
			switch err.(type) {
			case *adaptors.NoFaceError:
				br.ErrorType = models.NoFace
			case *adaptors.MultipleFaces:
				br.ErrorType = models.MultipleRecognized
			}
			log.LogBadRequest(br)
		}()
		return
	}
	profileId, err = ps.CreateProfile(name, image, descriptor, shop)
	return
}
