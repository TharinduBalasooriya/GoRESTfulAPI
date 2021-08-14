package db

import "github.com/TharinduBalasooriya/goRestApi/models"

var MovieList = []models.Movie{

	{
		Id:         "1",
		Title:      "The Godfather",
		Year:       "1972",
		Category:   "Cime/Drama",
		ImdbRating: "9.2/10",
	},
	{
		Id:         "2",
		Title:      "The English Patient",
		Year:       "1996",
		Category:   "Romance/Drama",
		ImdbRating: "7.4/10",
	}, {
		Id:         "3",
		Title:      "The Greate Gatsby",
		Year:       "2013",
		Category:   "Romance/Drama",
		ImdbRating: "7.2/10",
	},
}
