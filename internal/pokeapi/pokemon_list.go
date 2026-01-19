package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListPokemon -
func (c *Client) ListPokemon(pageURL *string, location string) (RespShallowPokemonList, error) {
	url := baseURL + "/location-area" + "/" + location
	if pageURL != nil {
		url = *pageURL
	}

	// Use cache instead of new request
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := RespShallowPokemonList{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespShallowPokemonList{}, err
		}

		return pokemonResp, nil
	}

	// make new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowPokemonList{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowPokemonList{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemonList{}, err
	}

	pokemonResp := RespShallowPokemonList{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespShallowPokemonList{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
