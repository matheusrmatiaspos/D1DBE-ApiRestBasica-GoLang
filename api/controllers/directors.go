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

var directorsDB = data.Directors

func GetDirectors(w http.ResponseWriter, r *http.Request){
	utils.ConsoleLog(w,r)
    w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(directorsDB)
}

func GetDirectorById(w http.ResponseWriter, r *http.Request){
	utils.ConsoleLog(w,r)
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range directorsDB{
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&models.Director{})
}

func CreateDirector(w http.ResponseWriter, r *http.Request){
	data,_ :=io.ReadAll(r.Body)

	var director models.Director

	err := json.Unmarshal(data,&director)
	utils.ConsoleLog(w,r)

	log.Printf("Diretor Inserido: %s",data)

	if err!=nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	directorsDB = append(directorsDB, director)
	json.NewEncoder(w).Encode(&director)
}