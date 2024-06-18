package main

func main() {

	cfg := &config{
		next: "https://pokeapi.co/api/v2/location",
		prev: "",
	}

	start(cfg)
}
