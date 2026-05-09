package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) Pokemon(pokemonName *string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + *pokemonName

	if cacheEntry, ok := c.cache.Get(url); ok {
		locationsResp := RespPokemon{}
		err := json.Unmarshal(cacheEntry, &locationsResp)
		if err != nil {
			return RespPokemon{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return RespPokemon{}, fmt.Errorf("Couldn't find the pokemon: %s Server says: %v", *pokemonName, resp.Status)
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	locationsResp := RespPokemon{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)

	return locationsResp, nil
}
