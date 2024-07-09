package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang/api/data"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang/api/models"
	"github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang/api/utils"
)

var moviesDB = data.Movies

func GetMovies(w http.ResponseWriter, r *http.Request) {
	utils.ConsoleLog(w,r)
    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moviesDB)
}

func GetMovieById(w http.ResponseWriter, r *http.Request){
	utils.ConsoleLog(w,r)
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range moviesDB{
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&models.Movie{})
}

func CreateMovie(w http.ResponseWriter, r *http.Request){
	data,_ :=io.ReadAll(r.Body)

	var movie models.Movie

	err := json.Unmarshal(data,&movie)
	utils.ConsoleLog(w,r)

	log.Printf("Filme Inserido: %s",data)

	if err!=nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	moviesDB = append(moviesDB, movie)
	json.NewEncoder(w).Encode(&movie)
}