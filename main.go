package main

import (
	"fmt"
	"github.com/patricekahlhorn/missing_episode_finder/internal/files"
	"github.com/patricekahlhorn/missing_episode_finder/internal/strings"
	"github.com/patricekahlhorn/missing_episode_finder/internal/tmdb"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

func main() {
	showName := ShowName()

	existingEpisodes := files.ExistingEpisodes()

	seasons := tmdb.GetSeasons(showName)

	wg := sync.WaitGroup{}
	wg.Add(len(seasons))

	for _, season := range seasons {
		go checkMissingEpisode(season, existingEpisodes, &wg)

	}

	wg.Wait()
}

type Episode struct {
	id     int
	season int
}

func (e Episode) number() string {

	number := strconv.Itoa(e.id)
	if len(number) <= 1 {
		number = "0" + number
	}

	return number
}

func checkMissingEpisode(s tmdb.Season, existingEpisodes []string, wg *sync.WaitGroup) {
	seasonNumber := s.Number()

	defer wg.Done()

	if seasonNumber == "" {
		return
	}

	for i := 1; i < s.EpisodeCount; i++ {
		episode := Episode{id: i}

		needle := strings.EpisodeString(seasonNumber, episode.number())

		exists := strings.ContainsSubstring(existingEpisodes, needle)

		if !exists {
			fmt.Println("Missing Episode " + needle)
		}
	}
}

func ShowName() string {
	wd, err := os.Getwd()

	if err != nil {
		panic("Error getting current working directory")
	}

	return filepath.Base(wd)
}
