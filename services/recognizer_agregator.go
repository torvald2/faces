package services

import (
	"strconv"
	"strings"
	"sync"

	"atbmarket.comfaceapp/config"
)

type RecognizeAgg struct {
	mu           sync.Mutex
	recognizers  map[int]FaceRecognizer
	profileStore ProfileStore
}

func (r RecognizeAgg) GetRecognizer(shopId int) (FaceRecognizer, bool) {
	recognizer, ok := r.recognizers[shopId]
	return recognizer, ok
}

func (r RecognizeAgg) ReinitRecognizer(shopId int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	profiles, err := r.profileStore.GetShopProfiles(shopId)
	if err != nil {
		return err
	}
	recignizer, err := NewRecognizer(profiles, shopId)
	if err != nil {
		return err
	}
	r.recognizers[shopId] = recignizer
	return nil
}

var recognizers RecognizeAgg

func CreateRecognizers(profileStore ProfileStore) (err error) {
	var rec FaceRecognizer
	conf := config.GetConfig()
	shops := strings.Split(conf.Shops, ",")
	recognizers.profileStore = profileStore
	recognizers.recognizers = make(map[int]FaceRecognizer)
	if len(shops) == 0 {
		panic("No shops in ACTIVE_SHOPS environment variable")
	}
	recognizers.mu.Lock()
	defer recognizers.mu.Unlock()
	for _, shop := range shops {
		shopNum, err := strconv.Atoi(shop)
		if err != nil {
			return err
		}
		profiles, err := recognizers.profileStore.GetShopProfiles(shopNum)
		if err != nil {
			continue
		}
		if conf.RecognizionMethod == "CUSTOM" {
			rec, err = NewCustomRecognizer(profiles, shopNum)
		} else {
			rec, err = NewRecognizer(profiles, shopNum)
		}
		if err != nil {
			return err
		}
		recognizers.recognizers[shopNum] = rec

	}
	return
}
