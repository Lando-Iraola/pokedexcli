package main

import "strings"

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	parts := strings.Split(text, " ")
	var in []string
	for _, part := range parts {
		if part != " " && part != "" {
			in = append(in, part)
		}
	}
	return in
}
