package services

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type RecognizeAgg struct {
	mu           sync.Mutex
	recognizers  map[int]Recognizer
	profileStore ProfileStore
}

func (r RecognizeAgg) GetRecognizer(shopId int) (Recognizer, bool) {
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
	shops := strings.Split(os.Getenv("ACTIVE_SHOPS"), ",")
	recognizers.profileStore = profileStore
	recognizers.recognizers = make(map[int]Recognizer)
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
		rec, err := NewRecognizer(profiles, shopNum)
		if err != nil {
			return err
		}
		recognizers.recognizers[shopNum] = rec

	}
	return
}
