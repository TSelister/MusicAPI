package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type song struct {
	Name   string `json:"name,omitempty"`
	Album  string `json:"album,omitempty"`
	Year   string `json:"year,omitempty"`
	Singer string `json:"singer,omitempty"`
}

var playlist []song

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/song", createSong).Methods("POST")
	r.HandleFunc("/song/", getPlaylist).Methods("GET")
	r.HandleFunc("/song/{name}", getSong).Methods("GET")
	r.HandleFunc("/song/{name}", putSong).Methods("PUT")
	r.HandleFunc("/song/{name}", deleteSong).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func getPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlist)
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range playlist {
		if item.Name == params["name"] {
			playlist = append(playlist[:index], playlist[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(playlist)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range playlist {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&song{})
}

func createSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var s song
	_ = json.NewDecoder(r.Body).Decode(&s)

	for _, item := range playlist {
		if item.Name == s.Name {
			w.WriteHeader(400)
			w.Write([]byte("song already exist"))
			return
		}
	}

	playlist = append(playlist, s)
	json.NewEncoder(w).Encode(playlist)
}

func putSong(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range playlist {
		if item.Name == params["name"] {
			playlist = append(playlist[:index], playlist[index+1:]...)
			var s song
			_ = json.NewDecoder(r.Body).Decode(&s)
			s.Name = params["name"]
			playlist = append(playlist, s)
			break
		}
	}
	json.NewEncoder(w).Encode(playlist)
}
