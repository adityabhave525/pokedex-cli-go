package main

import (
	"adityabhave525/pokedexcli/internal/pokeapi"
	"time"
)

func main() {

	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeApiClient: pokeClient,
	}

	startRepl(cfg)
}
