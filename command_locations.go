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

func fetchLocale(url string) (*Locales, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var locales Locales
	err = json.Unmarshal(body, &locales)
	if err != nil {
		return nil, err
	}

	return &locales, nil

}

func commandMap(cfg *config) error {
	next := cfg.next

	data, err := fetchLocale(next)
	if err != nil {
		return errors.New("no more results")
	}

	cfg.next = data.Next
	cfg.prev = data.Prev

	for i, id := range data.Results {
		fmt.Printf("%v: %s\n", i+1, id.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	prev := cfg.prev

	data, err := fetchLocale(prev)
	if err != nil {
		return errors.New("no more previous results")
	}
	fmt.Println(data.Prev)

	cfg.next = data.Next
	cfg.prev = data.Prev

	for i, id := range data.Results {
		fmt.Printf("%v: %s\n", i+1, id.Name)
	}

	return nil
}
