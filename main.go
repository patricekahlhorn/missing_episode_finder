package main

import (
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	wd, _ := os.Getwd()

	err := godotenv.Load("/home/server/projects/missing_episode/.env")
	if err != nil {
		log.Fatalf("Error loading %s .env file", wd)
	}

	tmdbClient, err := tmdb.InitV4(os.Getenv("TMDB_BEARER_TOKEN"))
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}

	basePath := filepath.Base(wd)

	search, err := tmdbClient.GetSearchTVShow(basePath, nil)

	if len(search.SearchTVShowsResults.Results) == 0 {
		panic("No TV shows found")
	}

	firstResult := search.SearchTVShowsResults.Results[0]

	tvShow, err := tmdbClient.GetTVDetails(int(firstResult.ID), nil)

	if err != nil {
		fmt.Println(err)
	}

	existingEpisodes := make([]string, 0)

	err = filepath.Walk(wd,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}

			existingEpisodes = append(existingEpisodes, info.Name())

			return nil
		})

	_ = existingEpisodes
	_ = tvShow

	for _, season := range tvShow.Seasons {
		//	fmt.Println(season.Name)
		//	fmt.Println(season.EpisodeCount)

		re := regexp.MustCompile("[0-9]+")
		seasonNumber := re.FindAllString(season.Name, -1)

		if seasonNumber == nil || len(seasonNumber) == 0 {
			continue
		}
		if len(seasonNumber[0]) <= 1 {
			seasonNumber[0] = "0" + seasonNumber[0]
		}

		for i := 1; i < season.EpisodeCount; i++ {

			episodeNumber := strconv.Itoa(i)
			if len(episodeNumber) <= 1 {
				episodeNumber = "0" + episodeNumber
			}

			needle := fmt.Sprintf("S%sE%s", seasonNumber[0], episodeNumber)

			blub := containsSubstring(existingEpisodes, needle)
			_ = blub
			if !blub {
				fmt.Println("Missing Episode " + needle)
			}
		}
	}

}

func containsSubstring(arr []string, substring string) bool {
	// Sort the array
	sort.Strings(arr)

	for _, v := range arr {
		if strings.Contains(strings.ToLower(v), strings.ToLower(substring)) {
			return true
		}
	}

	return false
}
