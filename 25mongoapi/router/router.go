package router

import (
	"github.com/gorilla/mux"
	"github.com/yagyagoel1/learning_golang/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("POST")
	router.HandleFunc("/api/deleteallmovie", controller.DeleteAllMovie).Methods("POST")

	return router
}
