package movies 

import (
	"net/http"
	"encoding/json"
)

const (
	omdbURL = "http://www.omdbapi.com/?i="
)

// MovieData contains the info returned from calls to OMDBApi.
type MovieData struct {
	Title      string `json:"Title"`
	Year	string `json:"Year"`
	Rated string `json:"Rated"`
	Released string `json:"Released"`
	Season string `json:"Season"`
	Episode string `json:"Episode"`
	Runtime string `json:"Runtime"`
	Genre string `json:"Genre"`
	Director string `json:"Director"`
	Writer string `json:"Writer"`
	Actors string `json:"Actors"`
	Plot string `json:"Plot"`
	Language string `json:"Language"`
	Country string `json:"Country"`
	Awards string `json:"Awards"`
	Poster string `json:"Poster"`
	Metascore string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes string `json:"imdbVotes"`
	ImdbID string `json:"imdbID"`
	SeriesID string `json:"seriesID"`
	Type string `json:"Type"`
	Response string `json:"Response"`
}

// GetMovieByID makes a request to OMDB and returns a movie using the imdb ID.
func GetMovieByID(imdbID string) (MovieData, error) {
	resp, err := http.Get(omdbURL + imdbID)

	if err != nil {
		return MovieData{}, err
	}

	defer resp.Body.Close()
	var m MovieData
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return MovieData{}, err
	}

	return m, nil
}