package models

type Actor struct{
	ID        string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	BirthYear  int `json:"birth_year,omitempty"`
	Nationality     string `json:"nationality,omitempty"`
	MoviesStarred   int `json:"movies_starred,omitempty"`
}