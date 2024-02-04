package handlers

import (
	"atbmarket.comfaceapp/services"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	//store := adaptors.GetDB()
	desc := services.NewDDescriptor()

	// err = services.CreateRecognizers(store)
	// if err != nil {
	// 	panic(fmt.Sprintf("Create recognize agregator error, %v", err))
	// }

	recognizeHandler := DiscHandler(desc)
	// newProfileHandler := GetNewFaceHandler(store)
	// imageHandler := GetImageHandler(store)
	// jornalHandler := GetWorkJornalHandler(store, store)
	// profileListHandler := GetProfileHandler(store)
	// getWorkJornal := GetSendWorkJornalHandler(store, adaptors.SendReport, adaptors.CreateSheet)
	// badRequestsHandler := GetBadRequestHandler(store)
	// distanceHandler := GetDistanceHandler(store)
	// badImageHandler := GetBadRequestImageHandler(store)

	// r.Use(getReqIdMidelware)

	api := r.PathPrefix("/api/").Subrouter()

	api.Handle("/dist", recognizeHandler).Methods("POST")
	// api.Handle("/profile/{id}", recognizeHandler).Methods("POST")
	// api.Handle("/profile/{id}/new", newProfileHandler).Methods("PUT")

	// api.Handle("/jornal", getWorkJornal).Methods("GET")
	// api.Handle("/jornal", jornalHandler).Methods("POST")

	// api.Handle("/badrequest", badRequestsHandler).Methods("GET")

	// api.Handle("/distance", distanceHandler).Methods("GET")

	// r.Handle("/images/{id}", imageHandler).Methods("GET")
	// r.Handle("/imagesrequest/{id}", badImageHandler).Methods("GET")

	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return r

}
