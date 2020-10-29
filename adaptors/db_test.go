package adaptors

import (
	"testing"

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

func TestGetAllProfiles(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Errorf("failed to load env")
	}
	store := GetDB()
	profiles, err := store.GetAllProfiles()
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
