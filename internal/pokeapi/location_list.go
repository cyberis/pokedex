package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocationAreaList(pageURL *string) (LocationAreaList, error) {
	url := baseURL + "/location-area"
	if pageURL != nil && *pageURL != "" {
		url = *pageURL
	}
	// Make the HTTP GET request and do many checks for errors
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request for location area list: %v", err)
		return LocationAreaList{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("Error fetching location area list: %v", err)
		return LocationAreaList{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return LocationAreaList{}, err
	}

	// Decode the JSON response into the LocationAreaList struct
	locationAreaList := LocationAreaList{}
	if err := json.Unmarshal(dat, &locationAreaList); err != nil {
		log.Printf("Error decoding location area list JSON: %v", err)
		return LocationAreaList{}, err
	}

	return locationAreaList, nil
}
