package services

import (
	"math"
	"strconv"

	"go.uber.org/zap"

	log "atbmarket.comfaceapp/app_logger"
	"atbmarket.comfaceapp/config"
	"atbmarket.comfaceapp/models"
	fr "github.com/Kagami/go-face"
	"gonum.org/v1/gonum/mat"
)

type CustomRecognizer struct {
	rec       *fr.Recognizer
	samples   map[int]*mat.VecDense
	shopId    int
	tolerance float32
}

func (r CustomRecognizer) GetUserIDByFace(image []byte) (userId int, err error) {
	face, err := r.getFace(image)
	if err != nil {
		return
	}

	userId = r.classify(face)
	if userId < 0 {
		return 0, UserNotFound{}
	}
	return
}

func (r CustomRecognizer) classify(userVector *mat.VecDense) int {
	for k, v := range r.samples {
		w := mat.NewVecDense(128, nil)
		w.SubVec(v, userVector)
		d := mat.Dot(w, w)
		dist := math.Sqrt(d)
		log.Logger.Debug("Distance", zap.Float64("Distance", dist))
		if float32(dist) < r.tolerance {
			return k
		}

	}
	return -1

}

func (r CustomRecognizer) getFace(image []byte) (vec *mat.VecDense, err error) {
	faces, err := r.rec.Recognize(image)
	if err != nil {
		return
	}
	if len(faces) == 0 {
		return vec, NoFaceError{}
	}
	if len(faces) > 1 {
		return vec, MultipleFaces{}
	}
	points := descriptorToFloatSlice(faces[0].Descriptor)
	vec = mat.NewVecDense(128, points)
	return
}

func (r CustomRecognizer) GetShopId() int {
	return r.shopId
}

func NewCustomRecognizer(data []models.Profile, shopId int) (rec CustomRecognizer, err error) {
	conf := config.GetConfig()
	samples := make(map[int]*mat.VecDense)
	//To-Do Use Descriptor
	r, err := fr.NewRecognizer("dnnModels")
	if err != nil {
		return
	}
	for _, v := range data {
		samples[v.Id] = mat.NewVecDense(128, v.Descriptor)
	}
	tolerance, err := strconv.ParseFloat(conf.Tolerance, 32)
	if err != nil {
		tolerance = 0.6
	}
	return CustomRecognizer{r, samples, shopId, float32(tolerance)}, nil
}
