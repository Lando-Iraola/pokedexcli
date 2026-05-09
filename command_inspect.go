package main

import "fmt"

func commandInspect(cfg *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("No pokemon name was given")
	}
	pokemon := params[0]

	if poke, ok := cfg.pokedex.Pokemon[pokemon]; !ok {
		return fmt.Errorf("you have not caught that pokemon")
	} else {
		fmt.Printf("Name: %s\n", poke.Name)
		fmt.Printf("Height: %d", poke.Height)
		fmt.Printf("Weight: %d", poke.Weight)
		fmt.Println("Stats:")
		for _, val := range poke.Stats {
			fmt.Printf("  -%s: %d\n", val.Stat.Name, val.BaseStat)
		}
		fmt.Println("Types:")
		for _, val := range poke.Types {
			fmt.Printf("  -%s\n", val.Type.Name)
		}
	}

	return nil
}
