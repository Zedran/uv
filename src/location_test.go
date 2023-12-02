package main

import (
	"testing"

	"github.com/Zedran/geoloc"
)

/*
	This test aims to check whether the specified string is properly formed into the Location struct by SpecifyLocation function.
	It does not check coordinate conversion.
*/
func TestSpecifyLocation(t *testing.T) {
	correctCases := map[string]geoloc.Location {
		"London, GB, 51.508, -0.128"     : {City: "London",  State: "", Country: "GB",  Lat: 51.508, Lon:  -0.128},
		", ,51.508,-0.128 "              : {City: "Unknown", State: "", Country: "N/A", Lat: 51.508, Lon:  -0.128},
		" Sendai , JP , 38.252, 140.856" : {City: "Sendai",  State: "", Country: "JP",  Lat: 38.252, Lon: 140.856},
	}

	incorrectCases := []string{
		"London, GB, 51,508, -0,128",
		"51.508, -0.128",
	}

	for tCase, val := range correctCases {
		out, err := SpecifyLocation(tCase)
		if err != nil || !(out.GetName(true) == val.GetName(true) && out.DistanceTo(&val) == 0) {
			t.Errorf(
				"Parsing failed:\n Got      : %s (%.3f, %.3f)\n Expected : %s (%.3f, %.3f)\n Err: %v",
				out.GetName(true), out.Lat, out.Lon, val.GetName(true), val.Lat, val.Lon, err,
			)
		}
	}

	for _, tCase := range incorrectCases {
		out, err := SpecifyLocation(tCase)
		if err == nil {
			t.Errorf(
				"Incorrect value did not trigger error:\n Got : {%s (%.3f, %.3f)} from \"%s\".",
				out.GetName(true), out.Lat, out.Lon, tCase,
			)
		}
	}
}
