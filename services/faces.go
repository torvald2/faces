package services

import (
	"fmt"

	"atbmarket.comfaceapp/adaptors"
	"atbmarket.comfaceapp/models"
)

func RecognizeFace(ra RecognizeAgregator, ps ProfileStore, image []byte, requestId string, shopId int) (profile models.Profile, err error) {
	recognozer, ok := ra.GetRecognizer(shopId)
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
			case *adaptors.NoFaceError:
				br.ErrorType = models.NoFace
			case *adaptors.MultipleFaces:
				br.ErrorType = models.MultipleRecognized
			case *adaptors.UserNotFound:
				br.ErrorType = models.UserNotFound

			}
			ps.LogBadRequest(br)
		}()
		return
	}
	profile, err = ps.GetProfileById(profileId)
	return
}

func CreateNewProfile(ra RecognizeAgregator, dg DescriptorGetter, ps ProfileStore, image []byte, name string, shop int, requestId string) (profileId int, err error) {

	descriptor, err := dg.GetNewFaceDescriptor(image)
	if err != nil {
		go func() {
			br := models.BadRequest{
				CurrentFace: image,
				RequestId:   requestId,
				Shop:        shop,
			}
			switch err.(type) {
			case *adaptors.NoFaceError:
				br.ErrorType = models.NoFace
			case *adaptors.MultipleFaces:
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
	err = ra.ReinitRecognizer(shop)

	return
}
