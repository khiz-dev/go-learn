package main

import "fmt"

func commandMap(state *config) error {
	next := state.state
	locations, next := getLocationAreas(next)
	if len(locations) == 0 {
		fmt.Println("No locations found.")
		return nil
	}
	for _, location := range locations {
		fmt.Println(location.Name)
	}

	if next != "" {
		state.setState(next)
	}
	return nil
}
