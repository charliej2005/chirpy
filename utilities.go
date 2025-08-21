package main

import "strings"

func cleanInput(body string, blocked map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		if _, ok := blocked[strings.ToLower(word)]; ok {
			words[i] = "****"
		}
	}
	filtered := strings.Join(words, " ")
	return filtered
}
