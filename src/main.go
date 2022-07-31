package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	log.SetFlags(0)

	originalWD, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(filepath.Dir(os.Args[0]))
	defer os.Chdir(originalWD)

	s, err := LoadSettings()
	if err != nil {
		log.Fatal(err)
	} else if s == nil {
		fmt.Println("Settings file has been generated.")
		return
	} else if len(s.OpenWeatherKey) == 0 {
		log.Fatal("OpenWeather API key not provided")
	} else if len(s.OpenUVKey) == 0 {
		log.Fatal("OpenUV API key not provided")
	}

	client := http.Client{
		Timeout: time.Minute,
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}

	loc, err := GetLocation(&client, s, "London,GB")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(loc.Name, loc.Country, loc.Lat, loc.Lon)
}
