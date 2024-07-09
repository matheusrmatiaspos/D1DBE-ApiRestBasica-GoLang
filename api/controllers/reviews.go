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

var reviewsDB = data.Reviews

func GetReviews(w http.ResponseWriter, r *http.Request){
	utils.ConsoleLog(w,r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviewsDB)
}

func GetReviewById(w http.ResponseWriter, r *http.Request){
	utils.ConsoleLog(w,r)
	w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range reviewsDB{
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&models.Review{})
}

func CreateReview(w http.ResponseWriter, r *http.Request){
	data,_ :=io.ReadAll(r.Body)

	var review models.Review

	err := json.Unmarshal(data,&review)
	utils.ConsoleLog(w,r)

	log.Printf("Análise Inserido: %s",data)

	if err!=nil {
		log.Fatalf("Erro ao converter JSON para struct: %s", err)
	}

	reviewsDB = append(reviewsDB, review)
	json.NewEncoder(w).Encode(&review)
}