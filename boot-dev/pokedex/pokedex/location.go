package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func getLocationAreas(next string) ([]location, string) {
	url := baseURL + "location-area"
	if next != "" {
		url = next
	}
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	body, err := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		log.Fatal("Error: ", res.Status)
	}
	responseBody := responseResult{}
	json.Unmarshal(body, &responseBody)
	return responseBody.Results, responseBody.Next
}
