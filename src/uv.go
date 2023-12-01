package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Zedran/geoloc"
)

// A template URL for requesting data from OpenUV API
const OPEN_UV_URL string = "https://api.openuv.io/api/v1/uv?lat=%f&lng=%f"

/* The top structure of the OpenUV API response. */
type UVReport struct {
	Result `json:"result"`
}

/* Result returned from OpenUV API call. */
type Result struct {
	UV        float32 `json:"uv"`
	UVMax     float32 `json:"uv_max"`
	Ozone     float32 `json:"ozone"`

	UVTime    string  `json:"uv_time"`
	UVMaxTime string  `json:"uv_max_time"`
	OzoneTime string  `json:"ozone_time"`

	SafeExposureTime  `json:"safe_exposure_time"`

	SunInfo           `json:"sun_info"`
}

/* Safe exposure time in minutes for different skin types (Fitzpatrick scale). */
type SafeExposureTime struct {
	ST1 int `json:"st1"` // very fair skin, white
	ST2 int `json:"st2"` // fair skin, white
	ST3 int `json:"st3"` // fair skin, cream white
	ST4 int `json:"st4"` // olive skin
	ST5 int `json:"st5"` // brown skin
	ST6 int `json:"st6"` // black skin
}

/* The top structure for the data related to the Sun. */
type SunInfo struct {
	SunTimes `json:"sun_times"`
}

/* Struct containing information on important sun positions. */
type SunTimes struct {
	Sunrise    string `json:"sunrise"`
	SolarNoon  string `json:"solarNoon"`
	Sunset     string `json:"sunset"`
	Night      string `json:"night"`

	GoldenHour string `json:"goldenHour"`
	GHMorning  string `json:"goldenHourEnd"`
}

/* Transforms the time values of UVReport struct into a more readable format. */
func (uv *UVReport) Reformat() {
	uv.UVTime     = ReformatTime(uv.UVTime,     HEADER_TIME_FORMAT)
	uv.UVMaxTime  = ReformatTime(uv.UVMaxTime,  DEFAULT_SUN_TIMES_FORMAT)
	uv.OzoneTime  = ReformatTime(uv.OzoneTime,  DEFAULT_SUN_TIMES_FORMAT)

	uv.Sunrise    = ReformatTime(uv.Sunrise,    DEFAULT_SUN_TIMES_FORMAT)
	uv.SolarNoon  = ReformatTime(uv.SolarNoon,  DEFAULT_SUN_TIMES_FORMAT)
	uv.Sunset     = ReformatTime(uv.Sunset,     DEFAULT_SUN_TIMES_FORMAT)
	uv.Night      = ReformatTime(uv.Night,      DEFAULT_SUN_TIMES_FORMAT)

	uv.GoldenHour = ReformatTime(uv.GoldenHour, DEFAULT_SUN_TIMES_FORMAT)
	uv.GHMorning  = ReformatTime(uv.GHMorning,  DEFAULT_SUN_TIMES_FORMAT)
}

/* Formats the UVReport struct into string. */
func (uv *UVReport) ToString() string {
	return fmt.Sprintf(
		"UV Index:\n"                 + 
		"  Current: %6.2f\n"          + 
		"  Max:     %6.2f (%s)\n"     + 
		"  Ozone:   %6.2f (%s)\n\n"        + 
		"Sunrise: %15s\n"             + 
		"Solar Noon: %12s\n"          + 
		"Sunset: %16s\n"              + 
		"Night: %17s\n"               + 
		"Golden Hour: %11s\n"         + 
		"Morning GH ends: %7s\n\n"    + 
		"Safe Exposure Time [min]:\n" + 
		"  1: %5d   |   4: %5d\n"     + 
		"  2: %5d   |   5: %5d\n"     + 
		"  3: %5d   |   6: %5d", 
		uv.UV, uv.UVMax, uv.UVMaxTime,
		uv.Ozone,
		uv.OzoneTime,
		uv.Sunrise,
		uv.SolarNoon,
		uv.Sunset,
		uv.Night, 
		uv.GoldenHour,
		uv.GHMorning,
		uv.ST1, uv.ST4, 
		uv.ST2, uv.ST5, 
		uv.ST3, uv.ST6,
	)
}

/* Requests the report from OpenUV API.*/
func GetUVReport(client *http.Client, loc *geoloc.Location, s *Settings) (*UVReport, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(OPEN_UV_URL, loc.Lat, loc.Lon), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("x-access-token", s.OpenUVKey)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var uvReport UVReport

	err = json.NewDecoder(resp.Body).Decode(&uvReport)

	uvReport.Reformat()
	
	return &uvReport, err
}
