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
	id, err := store.CreateProfile("Test", data, []float32{1.11, 2.222}, 1)
	if err != nil {
		t.Errorf("Db error occured: %v", err)
	}
	if id == 0 {
		t.Errorf("No profile id returned")
	}

}
