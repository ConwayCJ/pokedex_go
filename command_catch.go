package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide pokemon name")
	}

	pkReq := args[0]

	pokemon, err := fetchPokemon(pkReq, cfg)
	if err != nil {
		return errors.New("they escaped! try catching a different one")
	}

	// chance to catch
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	catchChance := r.Intn(240) + 30
	thresh := pokemon.BaseExperience

	fmt.Println("You throw a pokeball at " + pokemon.Name)
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
	fmt.Println()

	// Simulate the Pokéball shaking
	fmt.Println("The Pokéball is shaking...")
	for i := 0; i < 3; i++ {
		time.Sleep(700 * time.Millisecond)
		fmt.Printf("Ping...\n")
	}

	// Define the threshold for catching the Pokémon

	// Determine if the Pokémon is caught
	time.Sleep(700 * time.Millisecond)
	if catchChance >= thresh {
		fmt.Printf("Congratulations! You caught a wild %s\n", pokemon.Name)
		addToPokedex(pokemon, cfg)
	} else {
		fmt.Printf("Oh no! The wild %s escaped.\n", pokemon.Name)
	}

	return nil
}

func fetchPokemon(pokeName string, cfg *config) (Pokemon, error) {
	pokeUrl := "https://pokeapi.co/api/v2/pokemon/"
	url := pokeUrl + pokeName

	// if it exists in cache, return that
	if val, exists := cfg.cache.Get(url); exists {
		var pokemon Pokemon

		err := json.Unmarshal(val, &pokemon)
		if err != nil {
			return Pokemon{}, nil
		}

		return pokemon, nil
	}

	// doesn't exist in cache, make http req

	resp, err := http.Get(url)

	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	cfg.cache.Add(url, body)

	return pokemon, nil
}

func addToPokedex(pokemon Pokemon, cfg *config) error {

	// check if pokemon exists, if exists - replace it?
	poke, exists := cfg.pokedex[pokemon.Name]
	if exists {
		fmt.Println("You've already caught a " + poke.Name)
		return nil
	}

	// doesn't exist, so add to pokedex
	cfg.pokedex[pokemon.Name] = pokemon
	fmt.Println(pokemon.Name + " has been added to the pokedex!")

	return nil
}
