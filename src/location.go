package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

/* This struct represents a location, typically a city. */
type Location struct {
	// Name of a city
	Name    string  `json:"name"`

	// State in which the city is located
	State   string `json:"state"`

	// Country in which the city is located
	Country string  `json:"country"`

	// Coordinates (S and W are negative)
	Lat     float32 `json:"lat"`
	Lon     float32 `json:"lon"`
}

const (
	// Maximum number of matches returned from the Geocoding API
	MAX_RESP_LOCS     int    = 10

	// A template URL for querying the Geocoding API
	OPEN_WEATHER_URL  string = "https://api.openweathermap.org/geo/1.0/direct?q=%s&limit=%d&appid=%s"
)

// This error is returned by GetLocation if the Geocoding API does not provide any result for the query
var errLocationNotFound error = errors.New("location not found")

/* Queries the OpenWeather Geocoding API for the name specified by locName and returns a slice containing
 * matching location names.
 */
func GetLocations(client *http.Client, settings *Settings, locName string) ([]Location, error) {
	resp, err := client.Get(fmt.Sprintf(OPEN_WEATHER_URL, locName, MAX_RESP_LOCS, settings.OpenWeatherKey))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locations []Location

	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		stream, _ := io.ReadAll(resp.Body)
		fmt.Print(string(stream))
		return nil, err
	} else if len(locations) == 0 {
		return nil, errLocationNotFound
	}

	return locations, nil
}
