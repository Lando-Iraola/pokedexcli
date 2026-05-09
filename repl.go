package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Lando-Iraola/pokedexcli/internal/pokeapi"
)

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			txt := scanner.Text()
			txts := cleanInput(txt)
			if len(txts) == 0 {
				continue
			}

			params := txts[1:]

			if command, ok := getCommands()[txts[0]]; ok {
				err := command.callback(conf, params...)
				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Println("Unkown command")
				continue
			}
		}
	}
}

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       Dex
	Next          *string
	Previous      *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config, params ...string) error
}

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

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next pokemon locations",
			callback:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous pokemon locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display all pokemon available in a given area. Takes in the map name",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Try to capture a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Show details of your pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List your pokemon",
			callback:    commandPokedex,
		},
	}
}
