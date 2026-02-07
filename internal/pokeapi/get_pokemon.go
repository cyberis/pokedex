package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	// Check cache first
	if cachedData, found := c.cache.Get(url); found {
		pokemon := Pokemon{}
		if err := json.Unmarshal(cachedData, &pokemon); err != nil {
			log.Printf("Error decoding cached pokemon JSON: %v", err)
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	// Make the HTTP GET request and do many checks for errors
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request for pokemon: %v", err)
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("Error fetching pokemon: %v", err)
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return Pokemon{}, err
	}

	if string(dat) == "Not Found" {
		log.Printf("Pokemon %q not found", name)
		return Pokemon{}, err
	}

	// Decode the JSON response into the Pokemon struct
	pokemon := Pokemon{}
	if err := json.Unmarshal(dat, &pokemon); err != nil {
		log.Printf("Error decoding pokemon JSON: %v", err)
		log.Printf("Response body: %s", string(dat))
		return Pokemon{}, err
	}

	// Add the fetched data to the cache
	c.cache.Add(url, dat)

	return pokemon, nil
}

func (c *Client) GetPokemonBaseExperience(name string) (int, error) {
	pokemon, err := c.GetPokemon(name)
	if err != nil {
		return 0, err
	}
	return pokemon.BaseExperience, nil
}

func (c *Client) GetPokemonHeight(name string) (int, error) {
	pokemon, err := c.GetPokemon(name)
	if err != nil {
		return 0, err
	}
	return pokemon.Height, nil
}

func (c *Client) GetPokemonWeight(name string) (int, error) {
	pokemon, err := c.GetPokemon(name)
	if err != nil {
		return 0, err
	}
	return pokemon.Weight, nil
}

func (c *Client) GetPokemonStats(name string) (map[string]int, error) {
	pokemon, err := c.GetPokemon(name)
	if err != nil {
		return nil, err
	}
	stats := make(map[string]int)
	for _, stat := range pokemon.Stats {
		stats[stat.Stat.Name] = stat.BaseStat
	}
	return stats, nil
}

func (c *Client) GetPokemonTypes(name string) ([]string, error) {
	pokemon, err := c.GetPokemon(name)
	if err != nil {
		return nil, err
	}
	types := []string{}
	for _, t := range pokemon.Types {
		types = append(types, t.Type.Name)
	}
	return types, nil
}
