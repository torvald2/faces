package services

func GetProfiles(ps profileStore, shopId int) (profiles map[int]string, err error) {
	profiles = make(map[int]string)
	data, err := ps.GetShopProfiles(shopId)
	if err != nil {
		return
	}
	for _, profile := range data {
		profiles[profile.Id] = profile.Name
	}
	return
}
