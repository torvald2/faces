package handlers

import (
	"fmt"

	"atbmarket.comfaceapp/adaptors"
	"atbmarket.comfaceapp/services"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	store := adaptors.GetDB()
	err := services.NewDescriptor()
	if err != nil {
		panic("Create descriptor error")
	}
	err = services.CreateRecognizers(store)
	if err != nil {
		panic(fmt.Sprintf("Create recognize agregator error, %v", err))
	}

	recognizeHandler := GetRecognizeFaceHandler(store)
	newProfileHandler := GetNewFaceHandler(store)
	imageHandler := GetImageHandler(store)
	jornalHandler := GetWorkJornalHandler(store, store)
	profileListHandler := GetProfileHandler(store)
	getWorkJornal := GetSendWorkJornalHandler(store, adaptors.SendReport, adaptors.CreateSheet)
	badRequestsHandler := GetBadRequestHandler(store)

	r.Use(getReqIdMidelware)

	api := r.PathPrefix("/api/").Subrouter()

	api.Handle("/profile/{id}", profileListHandler).Methods("GET")
	api.Handle("/profile/{id}", recognizeHandler).Methods("POST")
	api.Handle("/profile/{id}/new", newProfileHandler).Methods("PUT")

	api.Handle("/jornal", getWorkJornal).Methods("GET")
	api.Handle("/jornal", jornalHandler).Methods("POST")

	api.Handle("/badrequest", badRequestsHandler).Methods("GET")

	r.Handle("/images/{id}", imageHandler).Methods("GET")

	return r

}
