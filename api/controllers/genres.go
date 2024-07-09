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

var genresDB = data.Genres

func GetGenres(w http.ResponseWriter, r *http.Request){
	utils.ConsoleLog(w,r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genresDB)
}

func GetGenreById(w http.ResponseWriter, r *http.Request){
	utils.ConsoleLog(w,r)
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range genresDB{
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&models.Genre{})
}

func CreateGenre(w http.ResponseWriter, r *http.Request){
	data,_ :=io.ReadAll(r.Body)

	var genre models.Genre

	err := json.Unmarshal(data,&genre)
	utils.ConsoleLog(w,r)

	log.Printf("Gênero Inserido: %s",data)

	if err!=nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	genresDB = append(genresDB, genre)
	json.NewEncoder(w).Encode(&genre)
}