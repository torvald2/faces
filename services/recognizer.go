package services

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
	rec    *fr.Recognizer
	shopId int
}

type descriptor struct {
	rec *fr.Recognizer
}

func (r Recognizer) GetUserIDByFace(image []byte) (userId int, err error) {
	face, err := r.getFace(image)
	if err != nil {
		return
	}

	userId = r.rec.ClassifyThreshold(face.Descriptor, 0.6)
	if userId < 0 {
		return 0, UserNotFound{}
	}
	return
}
func (d descriptor) GetNewFaceDescriptor(image []byte) (descriptor []float32, err error) {
	descriptor = make([]float32, 128)
	faces, err := d.rec.Recognize(image)
	if err != nil {
		return
	}
	if len(faces) == 0 {
		return descriptor, NoFaceError{}
	}
	if len(faces) > 1 {
		return descriptor, MultipleFaces{}
	}
	face := faces[0]
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

func (r Recognizer) GetShopId() int {
	return r.shopId
}

func NewRecognizer(data []models.Profile, shopId int) (rec Recognizer, err error) {
	var samples []face.Descriptor
	var avengers []int32

	tmp_rec, err := fr.NewRecognizer("dnnModels")
	if err != nil {
		return
	}
	for _, v := range data {
		avengers = append(avengers, int32(v.Id))
		samples = append(samples, floatSliceToDescriptor(v.Descriptor))
	}
	tmp_rec.SetSamples(samples, avengers)
	return Recognizer{tmp_rec, shopId}, nil
}

var desc descriptor

func NewDescriptor() (err error) {
	tmp_desc, err := fr.NewRecognizer("dnnModels")
	desc = descriptor{tmp_desc}
	return
}

func floatSliceToDescriptor(points []float64) fr.Descriptor {
	var dots fr.Descriptor
	for k, v := range points {
		dots[k] = float32(v)
	}
	return dots
}
