package main

import (
	"fmt"
)

func commandExplore(cfg *config) error {
	locationArea := cfg.arguments[0]
	resp, err := cfg.pokeapiClient.ListPokemonsInLocationArea(locationArea)
	if err != nil {
		return err
	}

	for _, pokemon := range resp.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}
