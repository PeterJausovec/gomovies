package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/mitchellh/colorstring"
	"github.com/peterj/gomovies/movies"
)

func main() {
	movie, err := movies.GetMovieByID(randomID())
	if err != nil {
		fmt.Println(colorstring.Color("[red]" + err.Error()))
	}

	textColor := getColorForRating(movie.ImdbRating)
	fmt.Println(colorstring.Color(textColor + movie.Title))
}

func getColorForRating(ratingStr string) string {
	textColor := "[default]"
	if ratingStr != "N/A" {
		rating, err := strconv.Atoi(ratingStr)

		if err != nil {
			rating = 1
		}

		if rating <= 5 {
			textColor = "[red]"
		} else if 5 < rating || rating > 7 {
			textColor = "[yellow]"
		} else if rating >= 7 {
			textColor = "[green]"
		} else {
			textColor = "[default]"
		}
	}

	return textColor
}

func randomID() string {
	rand.Seed(time.Now().UnixNano())
	res := "tt1"
	for i := 0; i < 6; i++ {
		res += strconv.Itoa(rand.Intn(9))
	}
	return res
}
