package main

import (
	"adityabhave525/pokedexcli/internal/pokeapi"
	"time"
)

func main() {

	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeApiClient: pokeClient,
	}

	startRepl(cfg)
}
