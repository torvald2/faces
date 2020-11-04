package handlers

import (
	"atbmarket.comfaceapp/adaptors"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	store := adaptors.GetDB()
	desc, err := adaptors.NewDescriptor()
	if err != nil {
		panic("Create descriptor error")
	}
	recognizers, err := adaptors.CreateRecognizers(store)
	if err != nil {
		panic("Create recognize agregator error")
	}

	recognizeHandler := GetRecognizeFaceHandler(store, recognizers)

	r.Handle("/{id}/recognize", recognizeHandler).Methods("POST")

	return r

}
