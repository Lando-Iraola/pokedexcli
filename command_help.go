package main

import "fmt"

func commandHelp(cfg *config) error {

	fmt.Println("\nWelcome to the Pokedex!\nUsage:")
	for _, command := range getCommands() {
		str := fmt.Sprintf("%s: %s", command.name, command.description)
		fmt.Println(str)
	}

	return nil

}
