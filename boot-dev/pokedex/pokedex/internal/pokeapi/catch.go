package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "pokemon/" + pokemonName

	// Checking Cache
	cacheRes, ok := c.cache.Get(url)
	if ok {
		pokemonsResp := Pokemon{}
		err := json.Unmarshal(cacheRes, &pokemonsResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonsResp, nil
	}

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
	// Adding to cache
	c.cache.Add(url, dat)

	pokemonsResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonsResp)
	if err != nil {
		return Pokemon{}, err
	}
	return pokemonsResp, nil
}
