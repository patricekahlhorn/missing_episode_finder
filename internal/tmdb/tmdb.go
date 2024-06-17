package tmdb

import (
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

type Season struct {
	AirDate      string  `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   string  `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float32 `json:"vote_average"`
}

func GetSeasons(basePath string) []Season {
	tmdbClient, err := tmdb.InitV4(token())

	search, err := tmdbClient.GetSearchTVShow(basePath, nil)

	if len(search.SearchTVShowsResults.Results) == 0 {
		panic("No TV shows found")
	}

	firstResult := search.SearchTVShowsResults.Results[0]

	tvShow, err := tmdbClient.GetTVDetails(int(firstResult.ID), nil)

	if err != nil {
		fmt.Println(err)
	}

	var seasons []Season

	for _, season := range tvShow.Seasons {
		seasons = append(seasons, season)
	}

	return seasons
}

func (s Season) Number() (Number string) {
	re := regexp.MustCompile("[0-9]+")
	seasonNumber := re.FindAllString(s.Name, -1)

	if seasonNumber == nil || len(seasonNumber) == 0 {
		return ""
	}
	if len(seasonNumber[0]) <= 1 {
		seasonNumber[0] = "0" + seasonNumber[0]
	}

	return PrependZero(seasonNumber[0])
}

func PrependZero(number string) string {
	if len(number) <= 1 {
		number = "0" + number
	}

	return number
}

func token() string {
	err := godotenv.Load("/home/server/projects/missing_episode/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	t := os.Getenv("TMDB_BEARER_TOKEN")

	if t == "" {
		panic("TMDB_BEARER_TOKEN environment variable not set")
	}

	return t
}
