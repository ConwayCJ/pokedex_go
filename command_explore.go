package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location")
	}

	area := args[0]

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon: ")
	locData, err := fetchLocation(area, cfg)
	if err != nil {
		return errors.New("no pokemon found")
	}

	for _, v := range locData.PokemonEncounters {
		fmt.Printf(v.Pokemon.Name)
	}
	return nil
}

func fetchLocation(reqArea string, cfg *config) (Location, error) {
	// check if in cache, if in cache - return
	url := "https://pokeapi.co/api/v2/location-area/" + reqArea
	if val, exists := cfg.cache.Get(url); exists {
		var loc Location

		err := json.Unmarshal(val, &loc)
		if err != nil {
			return Location{}, err
		}

		return loc, nil
	}

	// if not in cache, make http request - then add to cache
	resp, err := http.Get(url)

	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	var data Location
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Location{}, err
	}

	cfg.cache.Add(url, body)
	return data, nil
}
