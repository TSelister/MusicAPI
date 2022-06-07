package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type song struct {
	Name   string `json:"name,omitempty"`
	Album  string `json:"album,omitempty"`
	Year   string `json:"year,omitempty"`
	Singer string `json:"singer,omitempty"`
}

var database = make(map[string]song)

func main() {
	r := mux.NewRouter()
	r.Path("/song").Methods("POST").HandlerFunc(createSong)
	r.Path("/song/{name}").Methods("GET").HandlerFunc(getSong)
	r.Path("/song").Methods("PUT").HandlerFunc(putSong)
	r.Path("/song/{name}").Methods("DELETE").HandlerFunc(deleteSong)

	srv := http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadTimeout:       10,
		ReadHeaderTimeout: 10,
		WriteTimeout:      10,
	}

	srv.ListenAndServe()
}

func validateSong(s *song) error {
	if s.Name == "" {
		return errors.New("the name cannot be empty")
	}
	if s.Album == "" {
		return errors.New("the album cannot be empty")
	}
	if len(s.Year) < 3 {
		return errors.New("the year must have more than 3 characters")
	}
	if s.Singer == "" {
		return errors.New("the singer's name cannot be empty")
	}
	return nil
}

func createSong(w http.ResponseWriter, r *http.Request) {
	var s song

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = validateSong(&s)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	for _, item := range database {
		if item.Name == s.Name {
			w.WriteHeader(400)
			w.Write([]byte("song already exist"))
			return
		}
	}

	database[s.Name] = s

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(s)
}

func getSong(w http.ResponseWriter, r *http.Request) {
	q := mux.Vars(r)
	SongName := q["name"]

	song, ok := database[SongName]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Song not found"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(song)
}

func putSong(w http.ResponseWriter, r *http.Request) {
	var s song

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = json.Unmarshal(body, &s)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	_, ok := database[s.Name]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("song not found"))
		return
	}

	if s.Name == "" {
		w.WriteHeader(400)
		w.Write([]byte("the name cannot be empty"))
		return
	}

	err = validateSong(&s)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}

	database[s.Name] = s

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(body)
}

func deleteSong(w http.ResponseWriter, r *http.Request) {
	q := mux.Vars(r)
	SongName := q["name"]

	_, ok := database[SongName]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("song not found"))
		return
	}

	delete(database, SongName)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("music sucessfully deleted"))
}
