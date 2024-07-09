package models

type Genre struct{
	ID        string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	MovieCount     int `json:"movie_count,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}