package onepizza

import (
	"log"
	"strconv"
	"strings"
)

// TODO: Refactor string into io.Reader
func Parse(s string) Preferences {
	lines := strings.Split(s, "\n")
	lines = parseFirstLine(lines)

	var preferences Preferences
	for i := 0; i < len(lines); i = i + 2 {
		var nextPreference Preference
		likes := Ingredients{}.FromStrings(strings.Split(lines[i], " "))
		nextPreference.Likes = likes[1:]
		dislikes := Ingredients{}.FromStrings(strings.Split(lines[i+1], " "))
		nextPreference.Dislikes = dislikes[1:]
		preferences = append(preferences, nextPreference)
	}

	return preferences
}

func parseFirstLine(lines []string) []string {
	firstLine := lines[0]
	c, err := strconv.Atoi(firstLine)
	if err != nil {
		log.Fatalf("Unable to convert '%s' to integer: %v", firstLine, err)
	}
	lines = lines[1:]
	if len(lines) != 2*c {
		log.Fatalf("Invalid input, expected %d lines, but have %d", 2*c, len(lines))
	}
	return lines
}
