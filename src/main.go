package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Zedran/geoloc"
	"github.com/Zedran/uv/src/openuv"
)

var ROOT_DIR string

/* Displays locations matching the query and prompts the user to pick one. */
func ShowLocationPickingDialog(locations []geoloc.Location) *geoloc.Location {
	fmt.Print("Multiple locations were found.\n\n")

	for i, loc := range locations {
		fmt.Printf(
			"%3d:   %-40s   %7.3f, %8.3f\n", 
			i + 1, loc.GetName(true), loc.Lat, loc.Lon,
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

	flagFindLocation    := flag.String(
		"f", "", 
		"find location by name:\n" + 
		"  London       - just a city\n" + 
		"  \"London, GB\" - city and country (note the quotes and a comma)",
	)

	flagSpecifyCoords   := flag.String(
		"c", "",
		"similar to -l, specify comma-separated coordinates only:\n" +
		"  51.508,-0.128",
	)

	flagSpecifyLocation := flag.String(
		"l", "", 
		"specify own location:\n" +
		"  \"London, GB, 51.508, -0.128\" (comma-separated)",
	)

	flagDefaultLoc      := flag.Bool("d", false, "set the passed location as default")
	flagUnsetDefaultLoc := flag.Bool("u", false, "unsets the current default location")

	flag.Parse()

	var locFlagsActive int = 0

	if len(*flagFindLocation   ) > 0 { locFlagsActive++ }
	if len(*flagSpecifyCoords  ) > 0 { locFlagsActive++ }
	if len(*flagSpecifyLocation) > 0 { locFlagsActive++ }

	if locFlagsActive > 1 {
		log.Fatal("flags '-c', '-f' and '-l' cannot be used simultaneously")
	}

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

	var loc *geoloc.Location

	if len(*flagFindLocation) > 0 {
		if len(s.OpenWeatherKey) == 0 {
			log.Fatal("OpenWeather API key not provided")
		}

		locations, err := geoloc.FindLocation(&client, s.OpenWeatherKey, *flagFindLocation, geoloc.DEFAULT_MAX_RESP_LOCS)
		if err != nil {
			log.Fatal(err)
		}

		if len(locations) > 1 {
			loc = ShowLocationPickingDialog(locations)
		} else {
			loc = &locations[0]
		}
	} else if len(*flagSpecifyCoords) > 0 {
		loc, err = geoloc.SpecifyLocation(",," + *flagSpecifyCoords)
		if err != nil {
			log.Fatal(err)
		}
	} else if len(*flagSpecifyLocation) > 0 {
		loc, err = geoloc.SpecifyLocation(*flagSpecifyLocation)
		if err != nil {
			log.Fatal(err)
		}
	} else if s.DefaultLocation != nil {
		loc = s.DefaultLocation
	} else {
		log.Fatal("no location has been specified and the default one is not set")
	}

	if *flagDefaultLoc && locFlagsActive > 0 {
		s.DefaultLocation = loc
		if err = SaveSettings(s); err != nil {
			log.Fatal(err)
		}
	}

	uv, err := openuv.GetUVReport(&client, loc, s.OpenUVKey)
	if err != nil {
		log.Fatal(err)
	}

	ReformatReportTime(uv)

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

	fmt.Printf("%s    %.3f %.3f    %s    %s\n\n%s\n", 
		loc.GetName(false), loc.Lat, loc.Lon, uv.UVTime, reqString, ReportToString(uv))
}
