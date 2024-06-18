# PokeAPI Go Client

This Go project demonstrates how to make a GET request to the PokeAPI to retrieve data about various aspects of Pokemon.

## Project Description

The project fetches data from the PokeAPI, specifically details about a specified Pokémon. It performs an HTTP GET request to the PokeAPI, parses the JSON response, and prints the Pokémon's name, order, and types.

## Prerequisites

- Go 1.16 or higher
- Internet connection (to fetch data from the PokeAPI)

## Installation

1. **Clone the repository:**
    ```sh
    git clone https://github.com/conwaycj/pokedex_go.git
    cd pokedex_go
    ```

2. **Initialize Go modules (if not already initialized):**
    ```sh
    go mod init pokeapi-go-client
    ```

3. **Install dependencies (if any):**
    ```sh
    go mod tidy
    ```

## Usage

1. **Run the application:**
    ```sh
    go build && ./pokedex_go.exe
    ```

2. **See commands:**
    ```sh
    Pokedex > help
    ```

3. **`map` command example:**
    ```sh
    Pokedex > map
    ```

    The first-time output for this command will be:
    
    ```plaintext
    1: canalave-city
    2: eterna-city
    3: pastoria-city
    4: sunyshore-city
    5: sinnoh-pokemon-league
    ```