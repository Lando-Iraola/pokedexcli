package main

import "fmt"

func commandExplore(cfg *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("No map given")
	}
	mapName := "/" + params[0]
	details, err := cfg.pokeapiClient.LocationDetail(&mapName)
	if err != nil {
		return err
	}

	for _, detail := range details.PokemonEncounters {
		str := fmt.Sprintf("%s", detail.Pokemon.Name)
		fmt.Println(str)
	}

	return nil
}
