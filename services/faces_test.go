package services

import (
	"testing"

	"atbmarket.comfaceapp/models"
)

type storeMock struct{}

func (s storeMock) GetProfileById(int) (models.Profile, error)                        { return models.Profile{Id: 1}, nil }
func (s storeMock) CreateProfile(string, []byte, []float32, int) (int, error)         { return 1, nil }
func (s storeMock) GetShopProfiles(shopId int) (profiles []models.Profile, err error) { return }

type recognizerMock struct{}

func (r recognizerMock) GetUserIDByFace([]byte) (int, error) { return 1, nil }
func (r recognizerMock) GetNewFaceDescriptor([]byte) ([]float32, error) {
	return []float32{1.001, 2.002}, nil
}
func (r recognizerMock) GetShopId() int { return 1 }

type logMock struct{}

func (l logMock) LogBadRequest(request models.BadRequest) error { return nil }

var fr recognizerMock
var store storeMock
var log logMock

func TestRecognizeFace(t *testing.T) {
	profile, err := RecognizeFace(fr, store, log, []byte{1, 2, 3}, "123")
	if err != nil {
		t.Errorf("An error occured %v", err)
	}
	if profile.Id != 1 {
		t.Errorf("ProfileId must be 1, but given %v", profile.Id)
	}
}

func TestCreateNewProfile(t *testing.T) {
	profileId, err := CreateNewProfile(fr, store, log, []byte{1, 2, 3}, "Test", 1, "123")

	if err != nil {
		t.Errorf("An error occured %v", err)
	}
	if profileId != 1 {
		t.Errorf("ProfileId must be 1, but given %v", profileId)
	}

}
