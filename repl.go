package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var SUPPORTED_COMMANDS map[string]cliCommand

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	parts := strings.Split(text, " ")
	var in []string
	for _, part := range parts {
		if part != " " && part != "" {
			in = append(in, strings.ToLower(part))
		}
	}
	return in
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, command := range SUPPORTED_COMMANDS {
		fmt.Println(fmt.Sprintf("%s: %s", command.name, command.description))
	}

	return nil
}
