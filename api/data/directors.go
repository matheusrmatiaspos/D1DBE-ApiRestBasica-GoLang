package data

import "github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang/api/models"

var Directors = []models.Director{
	{
		ID: "1",
		Name: "Frank Darabont",
		BirthYear: 1959,
		Nationality: "Americano",
		MoviesDirected: 10,
	  },
	  {
		ID: "2",
		Name: "Francis Ford Coppola",
		BirthYear: 1939,
		Nationality: "Americano",
		MoviesDirected: 30,
	  },
	  {
		ID: "3",
		Name: "Christopher Nolan",
		BirthYear: 1970,
		Nationality: "Inglês",
		MoviesDirected: 12,
	  },
}