package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config) error {
	if len(cfg.caughtPokemons) == 0 {
		return errors.New("you have not caught any pokemons")
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}