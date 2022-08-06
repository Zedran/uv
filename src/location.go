package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

/* This struct represents a location, typically a city. */
type Location struct {
	// Name of a city
	Name    string  `json:"name"`

	// Country in which the city is located
	Country string  `json:"country"`

	// Coordinates (S and W are negative)
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

// A template URL for querying the Geocoding API
const OPEN_WEATHER_URL  string = "https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=%d&appid=%s"

// This error is returned by GetLocation if the Geocoding API does not provide any result for the query
var errLocationNotFound error = errors.New("location not found")

/* Queries the OpenWeather Geocoding API for the name specified by locName and returns Location struct
 * with the first result found. The name of the location should be as specific as possible
 * ('City, Country').
 */
func GetLocation(client *http.Client, settings *Settings, locName string) (*Location, error) {
	resp, err := client.Get(fmt.Sprintf(OPEN_WEATHER_URL, locName, 1, settings.OpenWeatherKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locations []Location

	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		return nil, err
	} else if len(locations) == 0 {
		return nil, errLocationNotFound
	}

	return &locations[0], nil
}
