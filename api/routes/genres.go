package routes

import (
	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang/api/controllers"
)

func Genres(router *mux.Router){
	router.HandleFunc("/generos",controllers.GetGenres).Methods("GET")
	router.HandleFunc("/gerero/{id}",controllers.GetGenreById).Methods("GET")
	router.HandleFunc("/generos",controllers.CreateGenre).Methods("POST")
}