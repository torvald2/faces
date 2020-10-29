package adaptors

import (
	"atbmarket.comfaceapp/models"
	"github.com/Kagami/go-face"
	fr "github.com/Kagami/go-face"
)

type NoFaceError struct{}
type MultipleFaces struct{}
type UserNotFound struct{}
type MultipleMatch struct{}

func (e NoFaceError) Error() string {
	return "Лицо не найдено"
}
func (e MultipleFaces) Error() string {
	return "Несколько лиц на фото"
}
func (e UserNotFound) Error() string {
	return "Пользователь не найден"
}
func (e MultipleMatch) Error() string {
	return "Совпадение с несколькими пользователями"
}

type Recognizer struct {
	rec *fr.Recognizer
}

func (r Recognizer) GetUserIDByFace(image []byte) (userId int, err error) {
	face, err := r.getFace(image)
	if err != nil {
		return
	}

	userId = r.rec.Classify(face.Descriptor)
	if userId < 0 {
		return 0, UserNotFound{}
	}
	return
}
func (r Recognizer) GetNewFaceDescriptor(image []byte) (descriptor []float32, err error) {
	descriptor = make([]float32, 128)
	face, err := r.getFace(image)
	if err != nil {
		return
	}
	desc := face.Descriptor
	for k, v := range desc {
		descriptor[k] = v
	}
	return
}

func (r Recognizer) getFace(image []byte) (fr.Face, error) {
	faces, err := r.rec.Recognize(image)
	if err != nil {
		return fr.Face{}, err
	}
	if len(faces) == 0 {
		return fr.Face{}, NoFaceError{}
	}
	if len(faces) > 1 {
		return fr.Face{}, MultipleFaces{}
	}
	return faces[0], nil
}

func NewRecognizer(data []models.Profile) (rec Recognizer, err error) {
	var samples []face.Descriptor
	var avengers []int32

	tmp_rec, err := fr.NewRecognizer("../dnnModels")
	if err != nil {
		return
	}
	for _, v := range data {
		avengers = append(avengers, int32(v.Id))
		samples = append(samples, floatSliceToDescriptor(v.Descriptor))
	}
	tmp_rec.SetSamples(samples, avengers)
	return Recognizer{tmp_rec}, nil
}

func floatSliceToDescriptor(points []float64) fr.Descriptor {
	var dots fr.Descriptor
	for k, v := range points {
		dots[k] = float32(v)
	}
	return dots
}
