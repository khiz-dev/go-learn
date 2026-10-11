package main

import (
	"time"

	"github.com/khiz-dev/pokedexcli/internal/pokeapi"
)

func main() {
	cfg := &config{
		commands:       getCommands(),
		pokeapiClient:  pokeapi.NewClient(5*time.Second, time.Minute*5),
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)
}
