package services

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

type recognizeAgg struct {
	mu               sync.Mutex
	recognizers      map[int]FaceRecognizer
	profileStore     ProfileStore
	recognizeCreator RecognizeCreator
}

func (r recognizeAgg) GetRecognizer(shopId int) FaceRecognizer {
	recognizer, _ := r.recognizers[shopId]
	return recognizer
}

func (r recognizeAgg) ReinitRecognizer(shopId int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	profiles, err := r.profileStore.GetShopProfiles(shopId)
	if err != nil {
		return err
	}
	recignizer, err := r.recognizeCreator(profiles, shopId)
	if err != nil {
		return err
	}
	r.recognizers[shopId] = recignizer
	return nil
}

var Recognizers recognizeAgg

func CreateRecognizers(profileStore ProfileStore, recognizeCreator RecognizeCreator) {
	shops := strings.Split(os.Getenv("ACTIVE_SHOPS"), ",")
	if len(shops) == 0 {
		panic("No shops in ACTIVE_SHOPS environment variable")
	}
	Recognizers.mu.Lock()
	defer Recognizers.mu.Unlock()
	for _, shop := range shops {
		shopNum, err := strconv.Atoi(shop)
		if err != nil {
			panic("Shop num must be integer")
		}
		profiles, err := Recognizers.profileStore.GetShopProfiles(shopNum)
		if err != nil {
			continue
		}
		rec, err := Recognizers.recognizeCreator(profiles, shopNum)
		if err != nil {
			panic(err)
		}
		Recognizers.recognizers[shopNum] = rec

	}
}
