package files

import (
	"os"
	"path/filepath"
)

func ExistingEpisodes() []string {
	existingEpisodes := make([]string, 0)

	filepath.Walk(".",
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

	return existingEpisodes
}
