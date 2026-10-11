package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config) error {
	if len(cfg.arguments) == 0 {
		return errors.New("you must provide a pokemon name")
	}

	pokemonName := cfg.arguments[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	resp, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	// Higher base experience means lower catch probability
	catchChance := 100.0 / (100.0 + float64(resp.BaseExperience))

	if rand.Float64() > catchChance {
		fmt.Printf("%s escaped!\n", resp.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", resp.Name)
	fmt.Println("You may now inspect it with the inspect command.")
	cfg.caughtPokemons[resp.Name] = resp
	return nil
}
