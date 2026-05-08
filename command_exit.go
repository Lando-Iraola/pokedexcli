package main

import (
	"fmt"
	"os"
)

func commandExit(c *config) func() error {
	cfg := c
	return func() error {
		if cfg == (&config{}) {
			fmt.Println("yeah...")
		}

		fmt.Println("Closing the Pokedex... Goodbye!")
		os.Exit(0)
		return nil
	}
}
