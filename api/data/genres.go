package data

import "github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang/api/models"

var Genres = []models.Genre{
	{
		ID: "1",
		Name: "Drama",
		Description: "Filmes de Drama",
		MovieCount: 500,
		CreatedAt: "2020-05-10",
	},
	{
		ID: "2",
		Name: "Crime",
		Description: "Filmes de crime",
		MovieCount: 200,
		CreatedAt: "2020-06-15",
	},
	{
		ID: "3",
		Name: "Ação",
		Description: "Filmes de Ação",
		MovieCount: 300,
		CreatedAt: "2020-07-20",
	},
}