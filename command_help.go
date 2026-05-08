package main

import "fmt"

func commandHelp(c *config) func() error {
	cfg := c
	return func() error {
		fmt.Println("\nWelcome to the Pokedex!\nUsage:")
		for _, command := range getCommands(cfg) {
			str := fmt.Sprintf("%s: %s", command.name, command.description)
			fmt.Println(str)
		}

		return nil
	}
}
