package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("No pokemon name was given")
	}

	pokemon := params[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	detail, err := cfg.pokeapiClient.Pokemon(&pokemon)
	if err != nil {
		return err
	}

	var str string
	chance := float64(detail.BaseExperience) - (rand.Float64() * float64(detail.BaseExperience))
	isCaught := chance > float64(detail.BaseExperience)*0.4
	if isCaught {
		cfg.pokedex.Pokemon[pokemon] = detail
		str = fmt.Sprintf("%s was caught!", detail.Name)

	} else {
		str = fmt.Sprintf("%s escaped!", detail.Name)
	}
	fmt.Println(str)
	return nil
}
