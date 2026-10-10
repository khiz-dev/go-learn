package pokeapi

import "net/http"

func (c *Client) GetLocationList(pageURL *string) (responseResult, error) {
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return responseResult{}, err
	}

	resp, err :- c.httpClient.Do(req)
	if err != nil {
		return responseResult{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return responseResult{}, err
	}
}
