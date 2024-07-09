package data

import "github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang/api/models"

var Movies = []models.Movie{
	{
		ID:          "1",
		Title:       "Som de Liberdade",
		Director:    "Frank Darabont",
		Genre:       "Drama",
		ReleaseYear: 1994,
	},
	{
		ID:          "2",
		Title:       "O Padrinho",
		Director:    "Francis Ford Coppola",
		Genre:       "Crime",
		ReleaseYear: 1972,
	},
	{
		ID:          "3",
		Title:       "Batman: O Cavaleiro das Trevas",
		Director:    "Christopher Nolan",
		Genre:       "Ação",
		ReleaseYear: 2008,
	},
}