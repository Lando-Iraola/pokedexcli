package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			txt := scanner.Text()
			txts := cleanInput(txt)
			if len(txts) == 0 {
				continue
			}

			if command, ok := getCommands()[txts[0]]; ok {
				err := command.callback()
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}
