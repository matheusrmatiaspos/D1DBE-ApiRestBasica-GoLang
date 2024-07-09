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

var actorsDB = data.Actors

func GetActors(w http.ResponseWriter, r *http.Request){
	utils.ConsoleLog(w,r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actorsDB)
}

func GetActorsByID(w http.ResponseWriter, r *http.Request){
	utils.ConsoleLog(w,r)
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range actorsDB{
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&models.Actor{})
}

func CreateActor(w http.ResponseWriter, r *http.Request){
	data,_ :=io.ReadAll(r.Body)

	var actor models.Actor

	err := json.Unmarshal(data,&actor)
	utils.ConsoleLog(w,r)

	log.Printf("Ator Inserido: %s",data)

	if err!=nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	actorsDB = append(actorsDB, actor)
	json.NewEncoder(w).Encode(&actor)
}