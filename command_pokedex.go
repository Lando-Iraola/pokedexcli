package main

import "fmt"

func commandPokedex(cfg *config, params ...string) error {

	fmt.Println("Your Pokedex:")

	for key, _ := range cfg.pokedex.Pokemon {
		fmt.Printf("  -%s\n", key)
	}

	return nil
}
