package main

import (
	"errors"
	"strings"

	"github.com/Zedran/geoloc"
)

// The error raised if the user-specified location (-l) is improperly structured
var errLocStringInvalid error = errors.New("improper structure of the specified location")

/* Converts the string specified with the -l flag into the Location struct. */
func SpecifyLocation(locString string) (*geoloc.Location, error) {
	separated := strings.Split(locString, ",")

	if len(separated) != 4 {
		return nil, errLocStringInvalid
	}

	for i := range separated {
		separated[i] = strings.TrimSpace(separated[i])
	}

	var (
		err error
		loc geoloc.Location
	)

	if len(separated[0]) == 0 {
		loc.City = "Unknown"
	} else {
		loc.City = separated[0]
	}

	if len(separated[1]) == 0 {
		loc.Country = "N/A"
	} else {
		loc.Country = separated[1]
	}
	
	loc.Lat, err = ConvertCoordinate(separated[2], 90)
	if err != nil {
		return nil, err
	}

	loc.Lon, err = ConvertCoordinate(separated[3], 180)
	if err != nil {
		return nil, err
	}

	return &loc, nil
}