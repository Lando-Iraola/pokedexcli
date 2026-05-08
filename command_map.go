package main

import (
	"fmt"
)

func commandMapF(cfg *config) error {

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.Next)

	if err != nil {
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Printf("jeh")
		str := fmt.Sprintf("%s", location.Name)
		fmt.Println(str)
	}

	return nil

}

func commandMapb(cfg *config) error {

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.Previous)

	if err != nil {
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, location := range locationResp.Results {
		fmt.Printf("jeh")
		str := fmt.Sprintf("%s", location.Name)
		fmt.Println(str)
	}

	return nil

}
