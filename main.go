package main

import (
	"time"

	pokecache "github.com/conwaycj/pokedex_go/internal"
)

// import pokecache "github.com/conwaycj/pokedex_go/internal"

func main() {

	cfg := &config{
		next:  "https://pokeapi.co/api/v2/location-area",
		prev:  "",
		cache: *pokecache.NewCache(5 * time.Minute),
	}

	start(cfg)
}
