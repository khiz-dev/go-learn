package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (responseResult, error) {
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Checking Cache
	cacheRes, ok := c.cache.Get(url)
	if ok {
		locationResp := responseResult{}
		err := json.Unmarshal(cacheRes, &locationResp)
		if err != nil {
			return responseResult{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return responseResult{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return responseResult{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseResult{}, err
	}
	// Adding to cache
	c.cache.Add(url, dat)

	locationResp := responseResult{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return responseResult{}, err
	}
	return locationResp, nil
}
