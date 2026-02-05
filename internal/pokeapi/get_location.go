package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	url := baseURL + "/location-area/" + name

	// Check cache first
	if cachedData, found := c.cache.Get(url); found {
		locationArea := LocationArea{}
		if err := json.Unmarshal(cachedData, &locationArea); err != nil {
			log.Printf("Error decoding cached location area JSON: %v", err)
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	// Make the HTTP GET request and do many checks for errors
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request for location area: %v", err)
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("Error fetching location area: %v", err)
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return LocationArea{}, err
	}

	if string(dat) == "Not Found" {
		log.Printf("Location area %q not found", name)
		return LocationArea{}, err
	}

	// Decode the JSON response into the LocationArea struct
	locationArea := LocationArea{}
	if err := json.Unmarshal(dat, &locationArea); err != nil {
		log.Printf("Error decoding location area JSON: %v", err)
		log.Printf("Response body: %s", string(dat))
		return LocationArea{}, err
	}

	// Add the fetched data to the cache
	c.cache.Add(url, dat)

	return locationArea, nil
}

func (c *Client) ListPokemonNamesInLocationArea(name string) ([]string, error) {
	locationArea, err := c.GetLocationArea(name)
	if err != nil {
		return nil, err
	}

	pokemonNames := []string{}
	for _, pokemonEnc := range locationArea.PokemonEncounters {
		pokemonNames = append(pokemonNames, pokemonEnc.Pokemon.Name)
	}

	return pokemonNames, nil
}
