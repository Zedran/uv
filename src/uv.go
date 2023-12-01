package main

import (
	"fmt"

	"github.com/Zedran/uv/src/openuv"
)

/* Transforms the time values of openuv.UVReport struct into a more readable format. */
func ReformatReportTime(uv *openuv.UVReport) {
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
func ReportToString(uv *openuv.UVReport) string {
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