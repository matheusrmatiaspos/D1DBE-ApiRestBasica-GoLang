package models

type Review struct{
	ID        string `json:"id,omitempty"`
	MovieId string `json:"movie_id,omitempty"`
	Reviewer  string `json:"reviewer,omitempty"`
	Rating     int `json:"rating,omitempty"`
	Comment   string `json:"comment,omitempty"`
}