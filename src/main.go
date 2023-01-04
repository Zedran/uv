package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var ROOT_DIR string

/* Displays locations matching the query and prompts the user to pick one. */
func ShowLocationPickingDialog(locations []Location) *Location {
	fmt.Println("Multiple locations were found.\n")

	for i, loc := range locations {
		fmt.Printf(
			"%3d:   %-15s%-26s%5s   %7.3f, %7.3f\n", 
			i + 1, loc.Name, loc.State, loc.Country, loc.Lat, loc.Lon,
		)
	}

	var n int

	fmt.Print("\nChoose a location: ")
	_, err := fmt.Scanln(&n)
	fmt.Println()

	if err != nil || n <= 0 || n > len(locations) {
		log.Fatal("invalid choice")
	}

	return &locations[n - 1]
}

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
		locations, err := GetLocations(&client, s, *flagLocation)
		if err != nil {
			log.Fatal(err)
		}

		if len(locations) > 1 {
			loc = ShowLocationPickingDialog(locations)
		} else {
			loc = &locations[0]
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

	reqString := ""

	// Setting request limit to -1 (or any negative) disables cache
	if s.RequestLimit > 0 {
		requestsLeft, err := ProcessRequestCounter(s)
		if err != nil {
			log.Fatal(err)
		}

		reqString = fmt.Sprintf("[%d]", requestsLeft)
	} else if s.RequestLimit == 0 {
		reqString = "[bad limit]" // 0 is not a proper request limit
	}

	fmt.Printf("%s, %s    %.3f %.3f    %s    %s\n\n%s\n", 
		loc.Name, loc.Country, loc.Lat, loc.Lon, uv.UVTime, reqString, uv.ToString())
}
