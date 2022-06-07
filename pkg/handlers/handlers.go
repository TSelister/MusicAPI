package handlers

import (
	"Estudos/API/pkg/mocks"
	"Estudos/API/pkg/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var counter int = 1

func CreateMusic(w http.ResponseWriter, r *http.Request) {
	var m models.Music

	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = json.Unmarshal(body, &m)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	for _, item := range mocks.Playlist {
		if item.Name == m.Name {
			w.WriteHeader(400)
			w.Write([]byte("song already exist"))
			return
		}
	}

	m.ID = counter
	counter++
	mocks.Playlist = append(mocks.Playlist, m)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}

func GetPlaylist(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Playlist)
}

func GetMusic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	for _, music := range mocks.Playlist {
		if music.Name == vars["name"] {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(music)
			break
		}
	}
}

func UpdateMusic(w http.ResponseWriter, r *http.Request) {
	var m models.Music
	vars := mux.Vars(r)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = json.Unmarshal(body, &m)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	for index, music := range mocks.Playlist {
		if music.Name == vars["name"] {
			music.Name = m.Name
			music.Album = m.Album
			music.Year = m.Year
			music.Singer = m.Singer

			mocks.Playlist[index] = music

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}

func DeleteMusic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for index, music := range mocks.Playlist {
		if music.ID == id {
			mocks.Playlist = append(mocks.Playlist[:index], mocks.Playlist[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}
