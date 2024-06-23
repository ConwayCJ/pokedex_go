package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Locales struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Prev    string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func fetchLocale(url string, cfg *config) (Locales, error) {
	// check if exists in cache, if exists - do something.

	if val, found := cfg.cache.Get(url); found {
		fmt.Println(">>>>>>using cached data<<<<<<")

		var locale Locales

		// unmarshal body
		if err := json.Unmarshal(val, &locale); err != nil {
			return Locales{}, fmt.Errorf("failed to unmarshal: %v", err)
		}

		cfg.next = locale.Next
		cfg.prev = locale.Prev

		return locale, nil

	}

	// by this point, doesn't exist in cache.
	// make http request & add to cache
	resp, err := http.Get(url)

	if err != nil {
		return Locales{}, nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locales{}, nil
	}

	var locales Locales
	err = json.Unmarshal(body, &locales)
	if err != nil {
		return Locales{}, nil
	}

	cfg.next = locales.Next
	cfg.prev = locales.Prev

	cfg.cache.Add(url, body)
	return locales, nil

}

func commandMap(cfg *config) error {
	next := cfg.next

	res, err := fetchLocale(next, cfg)
	if err != nil {
		return errors.New("no more results")
	}

	for i, id := range res.Results {
		fmt.Printf("%v: %s\n", i+1, id.Name)
	}

	cfg.next = res.Next
	cfg.prev = res.Prev

	return nil
}

func commandMapb(cfg *config) error {
	prev := cfg.prev

	if prev == "" {
		return errors.New("you're on the first page")
	}

	res, err := fetchLocale(prev, cfg)
	if err != nil {
		return errors.New("you're on the first page")
	}

	for i, id := range res.Results {
		fmt.Printf("%v: %s\n", i+1, id.Name)
	}

	cfg.next = res.Next
	cfg.prev = res.Prev

	return nil
}
