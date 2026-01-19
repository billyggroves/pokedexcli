package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemonInfo -
func (c *Client) GetPokemonInfo(pageURL *string, pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon
	if pageURL != nil {
		url = *pageURL
	}

	// Use cache instead of new request
	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}

		return pokemonResp, nil
	}

	// make new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
