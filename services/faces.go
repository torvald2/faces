package services

import (
	"atbmarket.comfaceapp/models"
)

func RecognizeFace(fr faceRecognizer, ps profileStore, log logger ,image []byte) (profile models.Profile, err error) {
	profileId, err := fr.GetUserIDByFace(image)
	if err != nil {
		return
	}
	profile, err = ps.GetProfileById(profileId)
	return

}

func CreateNewProfile(fr faceRecognizer, profileStore)