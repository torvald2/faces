package services

import (
	"time"

	"atbmarket.comfaceapp/models"
)

func RecognizeFace(ps ProfileStore, image []byte, requestId string, shopId int) (profile models.Profile, err error) {
	recognozer, ok := recognizers.GetRecognizer(shopId)
	err = NoProfilesForShop{}
	if !ok {
		return
	}
	profileId, err := recognozer.GetUserIDByFace(image)

	if err != nil {
		go func() {
			br := models.BadRequest{
				CurrentFace:   image,
				RequestId:     requestId,
				Shop:          shopId,
				RecognizeTime: time.Now(),
			}
			switch err.(type) {
			case NoFaceError:
				br.ErrorType = models.NoFace
				ps.LogBadRequest(br)
			case MultipleFaces:
				br.ErrorType = models.MultipleRecognized
				ps.LogBadRequest(br)
			case UserNotFound:
				br.ErrorType = models.UserNotFound
				ps.LogBadRequest(br)
			}
		}()
		return
	}
	profile, err = ps.GetProfileById(profileId)
	return
}

func CreateNewProfile(ps ProfileStore, image []byte, name string, shop int, requestId string) (profileId int, err error) {

	descriptor, err := desc.GetNewFaceDescriptor(image)
	if err != nil {
		go func() {
			br := models.BadRequest{
				CurrentFace:   image,
				RequestId:     requestId,
				Shop:          shop,
				RecognizeTime: time.Now(),
			}
			switch err.(type) {
			case NoFaceError:
				br.ErrorType = models.NoFace
				ps.LogBadRequest(br)
			case MultipleFaces:
				br.ErrorType = models.MultipleRecognized
				ps.LogBadRequest(br)
			}
		}()
		return
	}
	profileId, err = ps.CreateProfile(name, image, descriptor, shop)
	if err != nil {
		return
	}
	err = recognizers.ReinitRecognizer(shop)

	return
}
