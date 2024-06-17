package strings

import (
	"fmt"
	"sort"
	"strings"
)

func ContainsSubstring(arr []string, substring string) bool {
	// Sort the array
	sort.Strings(arr)

	for _, v := range arr {
		if strings.Contains(strings.ToLower(v), strings.ToLower(substring)) {
			return true
		}
	}

	return false
}

func EpisodeString(season string, episode string) string {
	return fmt.Sprintf("S%sE%s", season, episode)
}
