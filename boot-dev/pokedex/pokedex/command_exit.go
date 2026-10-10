package main

import (
	"fmt"
	"os"
)

func commandExit(state *config) error {
	state.setState("exit")
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
