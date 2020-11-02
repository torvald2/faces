package services

func GetImage(ig ImageGetter, profileId int) (image []byte, err error) {
	image, err = ig.GetImage(profileId)
	return
}
