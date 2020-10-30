package services

func GetImage(ig imageGetter, profileId int) (image []byte, err error) {
	image, err = ig.GetImage(profileId)
	return
}
