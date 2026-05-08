package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pokemap struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Location_response struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Pokemap `json:"results"`
}

func commandMap(c *config) func() error {
	cfg := c
	return func() error {
		if cfg.Next == "" && cfg.Previous != "" {
			return fmt.Errorf("There are no more locations")
		}

		if cfg.Next == "" && cfg.Previous == "" {
			cfg.Next = "https://pokeapi.co/api/v2/location-area"
		}

		resp, err := http.Get(c.Next)

		if err != nil {
			return fmt.Errorf("Failed to get the cities: %w", err)
		}

		if resp.StatusCode > 299 {
			return fmt.Errorf("Failed to get the cities: %v", resp.Status)
		}

		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		var lr Location_response
		err = decoder.Decode(&lr)

		if err != nil {
			return fmt.Errorf("Failed to parse json: %w", err)
		}

		cfg.Next = lr.Next
		cfg.Previous = lr.Previous

		for _, location := range lr.Results {
			fmt.Printf("jeh")
			str := fmt.Sprintf("%s", location.Name)
			fmt.Println(str)
		}

		return nil
	}
}

func commandMapBack(c *config) func() error {
	cfg := c
	return func() error {
		if cfg.Next == "" && cfg.Previous == "" || cfg.Previous == "" {
			return fmt.Errorf("There are no locations yet")
		}

		resp, err := http.Get(c.Previous)

		if err != nil {
			return fmt.Errorf("Failed to get the cities: %w", err)
		}

		if resp.StatusCode > 299 {
			return fmt.Errorf("Failed to get the cities: %v", resp.Status)
		}

		defer resp.Body.Close()
		decoder := json.NewDecoder(resp.Body)
		var lr Location_response
		err = decoder.Decode(&lr)

		if err != nil {
			return fmt.Errorf("Failed to parse json: %w", err)
		}

		cfg.Next = lr.Next
		cfg.Previous = lr.Previous
		for _, location := range lr.Results {
			fmt.Printf("jeh")
			str := fmt.Sprintf("%s", location.Name)
			fmt.Println(str)
		}

		return nil
	}
}
