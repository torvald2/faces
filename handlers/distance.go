package handlers

// func GetDistanceHandler(s services.ProfileStore) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		rid := getRequestID(r.Context())
// 		firstNumID, err := strconv.Atoi(r.FormValue("first_id"))
// 		if err != nil {
// 			responseWithError(err, w)
// 			log.Logger.Error("Profile ID error",
// 				zap.String("Method", r.Method),
// 				zap.String("URL", r.RequestURI),
// 				zap.String("RequestID", rid),
// 				zap.Error(err))
// 			return
// 		}
// 		secondNumID, err := strconv.Atoi(r.FormValue("second_id"))
// 		if err != nil {
// 			responseWithError(err, w)
// 			log.Logger.Error("Profile ID error",
// 				zap.String("Method", r.Method),
// 				zap.String("URL", r.RequestURI),
// 				zap.String("RequestID", rid),
// 				zap.Error(err))
// 			return
// 		}
// 		//distance, err := services.GetDistance(firstNumID, secondNumID, s)

// 		if err != nil {
// 			responseWithError(err, w)
// 			log.Logger.Error("Get Distance ERR",
// 				zap.String("Method", r.Method),
// 				zap.String("URL", r.RequestURI),
// 				zap.String("RequestID", rid),
// 				zap.Error(err))
// 			return
// 		}
// 		responseOk(w, map[string]float64{"distance": distance})

// 	})

// }
