package services

import (
	"math"
	"strconv"

	"github.com/Kagami/go-face"
	fr "github.com/Kagami/go-face"
	"github.com/torvald2/faces/config"
	"github.com/torvald2/faces/models"
	"gonum.org/v1/gonum/mat"
)

type NoFaceError struct{}
type MultipleFaces struct{}
type UserNotFound struct{}
type MultipleMatch struct{}
type NoProfilesForShop struct{}

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

func (e NoProfilesForShop) Error() string {
	return "Для данного магазина нет ни одного профиля"
}

type Recognizer struct {
	rec       *fr.Recognizer
	shopId    int
	tolerance float32
}

type descriptor struct {
	rec *fr.Recognizer
}

type Descriptor struct {
	rec *fr.Recognizer
}

func (r Recognizer) GetUserIDByFace(image []byte, requestId string) (userId int, err error) {
	face, err := r.getFace(image)
	if err != nil {
		return
	}

	userId = r.rec.ClassifyThreshold(face.Descriptor, r.tolerance)
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

func (d Descriptor) GetNewFaceDescriptor(image []byte) (descriptor []float64, err error) {
	descriptor = make([]float64, 128)
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
		descriptor[k] = float64(v)
	}
	return
}

func GetProfDistance(desc1, decf2 []float64) (d float64, err error) {

	v1 := mat.NewVecDense(128, desc1)
	v2 := mat.NewVecDense(128, decf2)
	w := mat.NewVecDense(128, nil)
	w.SubVec(v1, v2)
	disc := mat.Dot(w, w)
	d = math.Sqrt(disc)
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
	conf := config.GetConfig()
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
	tolerance, err := strconv.ParseFloat(conf.Tolerance, 32)
	if err != nil {
		tolerance = 0.6
	}
	return Recognizer{tmp_rec, shopId, float32(tolerance)}, nil
}

var desc descriptor

func NewDescriptor() (err error) {
	tmp_desc, err := fr.NewRecognizer("dnnModels")
	desc = descriptor{tmp_desc}
	return
}

func NewDDescriptor() Descriptor {
	tmp_desc, _ := fr.NewRecognizer("dnnModels")
	d := Descriptor{rec: tmp_desc}
	return d
}

func floatSliceToDescriptor(points []float64) fr.Descriptor {
	var dots fr.Descriptor
	for k, v := range points {
		dots[k] = float32(v)
	}
	return dots
}

func descriptorToFloatSlice(desc fr.Descriptor) []float64 {
	points := make([]float64, 128)
	for k, v := range desc {
		points[k] = float64(v)
	}
	return points
}
