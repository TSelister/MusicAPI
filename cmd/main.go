package main

import (
	"Estudos/API/pkg/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/music", handlers.CreateMusic).Methods("POST")
	r.HandleFunc("/music/", handlers.GetPlaylist).Methods("GET")
	r.HandleFunc("/music/{name}", handlers.GetMusic).Methods("GET")
	r.HandleFunc("/music/{name}", handlers.UpdateMusic).Methods("PUT")
	r.HandleFunc("/music/{name}", handlers.DeleteMusic).Methods("DELETE")

	log.Println("API is running!!")
	http.ListenAndServe(":8080", r)

}
