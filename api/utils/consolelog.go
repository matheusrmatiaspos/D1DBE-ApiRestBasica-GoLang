package utils

import (
	"io"
	"log"
	"net/http"
)

func ConsoleLog(w http.ResponseWriter, r *http.Request) {
	data,_ :=io.ReadAll(r.Body)
	log.Printf("%s%s - [%s] %s",r.Host,r.URL,r.Method, data)
}