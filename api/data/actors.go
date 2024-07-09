package data

import "github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang/api/models"

var Actors = []models.Actor{
	{
		ID: "1",
		Name: "Morgan Freeman",
		BirthYear: 1937,
		Nationality: "American",
		MoviesStarred: 100,
	  },
	  {
		ID: "2",
		Name: "Marlon Brando",
		BirthYear: 1924,
		Nationality: "American",
		MoviesStarred: 50,
	  },
	  {
		ID: "3",
		Name: "Christian Bale",
		BirthYear: 1974,
		Nationality: "British",
		MoviesStarred: 45,
	  },
}