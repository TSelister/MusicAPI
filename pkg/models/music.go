package models

type Music struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Album  string `json:"album"`
	Year   string `json:"year"`
	Singer string `json:"singer"`
}
