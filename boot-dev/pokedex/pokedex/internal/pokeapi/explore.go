package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemonsInLocationArea(locationArea string) (exploreResponseResult, error) {
	url := baseURL + "location-area/" + locationArea

	// Checking Cache
	cacheRes, ok := c.cache.Get(url)
	if ok {
		pokemonsResp := exploreResponseResult{}
		err := json.Unmarshal(cacheRes, &pokemonsResp)
		if err != nil {
			return exploreResponseResult{}, err
		}
		return pokemonsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return exploreResponseResult{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return exploreResponseResult{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return exploreResponseResult{}, err
	}
	// Adding to cache
	c.cache.Add(url, dat)

	pokemonsResp := exploreResponseResult{}
	err = json.Unmarshal(dat, &pokemonsResp)
	if err != nil {
		return exploreResponseResult{}, err
	}
	return pokemonsResp, nil
}
