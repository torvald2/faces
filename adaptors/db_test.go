package adaptors

import (
	"testing"
	"time"

	"github.com/torvald2/faces/models"

	"github.com/joho/godotenv"
)

func TestStore_CreateProfile(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("failed to load env")
	}
	store := GetDB()
	data := []byte{1, 2, 3}
	descriptor := make([]float32, 128)
	id, err := store.CreateProfile("Test", data, descriptor, 1)
	if err != nil {
		t.Errorf("Db error occured: %v", err)
	}
	if id == 0 {
		t.Errorf("No profile id returned")
	}

}

func TestGetShopProfiles(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("failed to load env")
	}
	store := GetDB()
	profiles, err := store.GetShopProfiles(1)
	if err != nil {
		t.Errorf("An error occured during load profiles %v", err)
	}
	if len(profiles) == 0 {
		t.Errorf("Function GetAllProfiles must return more then 0 profiles")
	}

}

func TestGetImage(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("failed to load env")
	}
	store := GetDB()
	data, err := store.GetImage(4)
	if err != nil {
		t.Errorf("Get image err %v", err)
	}
	if len(data) == 0 {
		t.Errorf("Data is empty")
	}

}

func TestNewJornalRecord(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("failed to load env")
	}
	store := GetDB()
	jornal := models.JornalOperation{
		UserId:        1,
		OperationDate: time.Now(),
		OperationType: "1",
	}
	if err := store.NewJornalRecord(jornal); err != nil {
		t.Errorf("Error in insert jornal operation %v", err)
	}
}

func TestLoadRequest(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("failed to load env")
	}
	store := GetDB()
	request := models.BadRequest{
		UserId:          1,
		RecognizedUsers: []int{1, 2, 3},
		CurrentFace:     []byte{1, 2, 3},
		RecognizeTime:   time.Now(),
		ErrorType:       "1",
		Shop:            1,
		RequestId:       "123",
	}
	if err := store.LogBadRequest(request); err != nil {
		t.Errorf("Error in insert bad request operation %v", err)
	}
}
