package data

import "github.com/matheurmatiaspos/D1DBE-ApiRestBasica-GoLang/api/models"

var Reviews = []models.Review{
	{
		ID: "1",
		MovieId: "1",
		Reviewer: "John Doe",
		Rating: 5,
		Comment: "Um filme incrível com uma história comovente.",
	},
	{
		ID: "2",
		MovieId: "2",
		Reviewer: "Jane Smith",
		Rating: 5,
		Comment: "Uma obra-prima clássica com performances excepcionais.",
	},
	{
		ID: "3",
		MovieId: "3",
		Reviewer: "Alice Johnson",
		Rating: 4,
		Comment: "Um filme emocionante com excelente direção.",
	},
}