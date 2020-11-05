package services

import (
	"testing"

	"atbmarket.comfaceapp/models"
)

func TestGetRecognizer(t *testing.T) {
	simple_profiles := make([]models.Profile, 100)
	_, err := newRecognizer(simple_profiles, 1)
	if err != nil {
		t.Errorf("Get recognizer error %v", err)
	}

}
