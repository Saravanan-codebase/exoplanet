package main

import (
	"log"
	"net/http"

	"exoplanet/handler"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/exoplanets", handler.AddExoplanet).Methods(http.MethodPost)
	router.HandleFunc("/exoplanets", handler.ListExoplanets).Methods(http.MethodGet)
	router.HandleFunc("/exoplanets/{id}", handler.GetExoplanetByID).Methods(http.MethodGet)
	router.HandleFunc("/exoplanets/{id}", handler.UpdateExoplanet).Methods(http.MethodPut)
	router.HandleFunc("/exoplanets/{id}", handler.DeleteExoplanet).Methods(http.MethodDelete)
	router.HandleFunc("/exoplanets/{id}/fuel-estimation", handler.CalculateFuelEstimation).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
