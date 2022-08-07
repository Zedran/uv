package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var ROOT_DIR string

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

	var err error

	ROOT_DIR, err = GetRootDir()
	if err != nil {
		log.Fatal(err)
	}

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

	uv, err := GetUVReport(&client, loc, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s, %s    %.3f %.3f    %s\n\n%s\n", loc.Name, loc.Country, loc.Lat, loc.Lon, uv.UVTime, uv.ToString())
}
