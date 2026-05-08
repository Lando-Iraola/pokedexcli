package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	SUPPORTED_COMMANDS = map[string]cliCommand{
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

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			txt := scanner.Text()
			txts := cleanInput(txt)

			if command, ok := SUPPORTED_COMMANDS[txts[0]]; ok {
				command.callback()
			}

		}
	}
}
