package models

type Director struct{
	ID        string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	BirthYear  int `json:"birth_year,omitempty"`
	Nationality     string `json:"nationality,omitempty"`
	MoviesDirected   int `json:"movies_directed,omitempty"`
}