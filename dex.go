package main

import "github.com/Lando-Iraola/pokedexcli/internal/pokeapi"

type Dex struct {
	Pokemon map[string]pokeapi.RespPokemon
}
