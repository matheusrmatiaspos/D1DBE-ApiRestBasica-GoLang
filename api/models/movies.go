package models

type Movie struct{
	ID        string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Director  string `json:"director,omitempty"`
	Genre     string `json:"genre,omitempty"`
	ReleaseYear   int `json:"release_year,omitempty"`
}