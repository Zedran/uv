package main

import (
	"encoding/json"
	"os"
	"path"
)

const (
	SETTINGS_DIR  string = "settings"
	SETTINGS_FILE string = SETTINGS_DIR + "/uv.json"
)

/* This struct represents the application's settings. */
type Settings struct {
	// OpenWeather API key for its geocoding service
	OpenWeatherKey  string    `json:"open_weather_key"`
	
	// OpenUV API key
	OpenUVKey       string    `json:"open_uv_key"`

	// Default location for which the report is generated if no other location is specified
	DefaultLocation *Location `json:"default_location"`

	// Request limit for user's OpenUV account. Setting it to -1 disables cache.
	RequestLimit    int64     `json:"request_limit"`
}

/* Loads settings from file. If the file is not present, it calls SaveSettings to generate
 * an empty one.
 */
func LoadSettings() (*Settings, error) {
	stream, err := os.ReadFile(path.Join(ROOT_DIR, SETTINGS_FILE))
	if err != nil {
		return nil, SaveSettings(nil)
	}

	var s Settings

	if err = json.Unmarshal(stream, &s); err != nil {
		return nil, err
	}

	return &s, nil
}

/* Saves settings to a file. If s is nil, the empty file is generated. */
func SaveSettings(s *Settings) error {
	if s == nil {
		s = &Settings{DefaultLocation: nil, RequestLimit: -1}
	}

	if _, err := os.Stat(path.Join(ROOT_DIR, SETTINGS_DIR)); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(ROOT_DIR, SETTINGS_DIR), 0775)
		if err != nil {
			return err
		}
	}

	stream, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		return err
	}

	f, err := os.OpenFile(path.Join(ROOT_DIR, SETTINGS_FILE), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0775)
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write(stream)

	return nil
}
