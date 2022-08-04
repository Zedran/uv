package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	log.SetFlags(0)

	flagLocation := flag.String("l", "", 
								"specify report's location:\n" + 
								"  London       - just a city\n" + 
								"  \"London, GB\" - city and country (note the quotes and a comma)")

	flagDefaultLoc      := flag.Bool("d", false, "set the passed location as default")
	flagUnsetDefaultLoc := flag.Bool("u", false, "unsets the current default location")

	flag.Parse()

	if *flagDefaultLoc && *flagUnsetDefaultLoc {
		log.Fatal("flags '-d' and '-u' cannot be used simultaneously")
	}

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

	if *flagUnsetDefaultLoc {
		s.DefaultLocation = nil

		if err = SaveSettings(s); err != nil {
			log.Fatal(err)
		}
		return
	}

	client := http.Client{
		Timeout: time.Minute,
		Transport: &http.Transport{
			DisableKeepAlives: true,
		},
	}

	var loc *Location

	if len(*flagLocation) > 0 {
		loc, err = GetLocation(&client, s, *flagLocation)
		if err != nil {
			log.Fatal(err)
		}

		if *flagDefaultLoc {
			s.DefaultLocation = loc
			if err = SaveSettings(s); err != nil {
				log.Fatal(err)
			}
		}
	} else if s.DefaultLocation != nil {
		loc = s.DefaultLocation
	} else {
		log.Fatal("no location has been specified and the default one is not set")

	}

	fmt.Println(loc.Name, loc.Country, loc.Lat, loc.Lon)
}
