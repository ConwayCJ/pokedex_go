package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a location")
	}

	pokeURL := "https://pokeapi.co/api/v2/location-area/"
	area := args[0]

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon: ")
	locData, err := fetchLocation(pokeURL+area, cfg)
	if err != nil {
		return errors.New("no pokemon found")
	}

	fmt.Println(locData.PokemonEncounters[0].Pokemon.Name)
	for k, v := range locData.PokemonEncounters {
		fmt.Printf("%v: %s\n", k, v.Pokemon.Name)
	}
	return nil
}

func fetchLocation(url string, cfg *config) (Location, error) {
	// check if in cache, if in cache - return
	fmt.Println(url)
	if val, exists := cfg.cache.Get(url); exists {
		var data Location

		err := json.Unmarshal(val, &data)
		if err != nil {
			return Location{}, err
		}

		return data, nil
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
	fmt.Println("returning data")
	return data, nil
}
