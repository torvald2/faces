package services

import (
	"fmt"

	"atbmarket.comfaceapp/models"
)

func RecognizeFace(ps ProfileStore, image []byte, requestId string, shopId int) (profile models.Profile, err error) {
	recognozer, ok := recognizers.GetRecognizer(shopId)
	err = fmt.Errorf("Для данного магазина нет ни одного профиля")
	if !ok {
		return
	}
	profileId, err := recognozer.GetUserIDByFace(image)

	if err != nil {
		go func() {
			br := models.BadRequest{
				CurrentFace: image,
				RequestId:   requestId,
				Shop:        shopId,
			}
			switch err.(type) {
			case *NoFaceError:
				br.ErrorType = models.NoFace
			case *MultipleFaces:
				br.ErrorType = models.MultipleRecognized
			case *UserNotFound:
				br.ErrorType = models.UserNotFound

			}
			ps.LogBadRequest(br)
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
				CurrentFace: image,
				RequestId:   requestId,
				Shop:        shop,
			}
			switch err.(type) {
			case *NoFaceError:
				br.ErrorType = models.NoFace
			case *MultipleFaces:
				br.ErrorType = models.MultipleRecognized
			}
			ps.LogBadRequest(br)
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
