package main

import (
	"strings"
)

func cleanInput(text string) []string {
	words := strings.Split(text, " ")
	var cleaned []string
	for _, word := range words {
		trimmed := strings.TrimSpace(word)
		lowered := strings.ToLower(trimmed)

		if lowered != "" {
			cleaned = append(cleaned, lowered)
		}
	}
	return cleaned
}
