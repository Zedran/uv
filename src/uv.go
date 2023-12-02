package main

import (
	"fmt"

	"github.com/Zedran/go-openuv"
)

/* Transforms the time values of openuv.UVReport struct into a more readable format. */
func ReformatReportTime(uv *openuv.UVReport) {
	st := &uv.SunInfo.SunTimes

	uv.UVTime        = ReformatTime(uv.UVTime,        HEADER_TIME_FORMAT)
	uv.UVMaxTime     = ReformatTime(uv.UVMaxTime,     DEFAULT_SUN_TIMES_FORMAT)
	uv.OzoneTime     = ReformatTime(uv.OzoneTime,     DEFAULT_SUN_TIMES_FORMAT)

	st.Sunrise       = ReformatTime(st.Sunrise,       DEFAULT_SUN_TIMES_FORMAT)
	st.SolarNoon     = ReformatTime(st.SolarNoon,     DEFAULT_SUN_TIMES_FORMAT)
	st.Sunset        = ReformatTime(st.Sunset,        DEFAULT_SUN_TIMES_FORMAT)
	st.Night         = ReformatTime(st.Night,         DEFAULT_SUN_TIMES_FORMAT)

	st.GoldenHour    = ReformatTime(st.GoldenHour,    DEFAULT_SUN_TIMES_FORMAT)
	st.GoldenHourEnd = ReformatTime(st.GoldenHourEnd, DEFAULT_SUN_TIMES_FORMAT)
}

/* Formats the UVReport struct into string. */
func ReportToString(uv *openuv.UVReport) string {
	return fmt.Sprintf(
		"UV Index:\n"                 + 
		"  Current: %6.2f\n"          + 
		"  Max:     %6.2f (%s)\n"     + 
		"  Ozone:   %6.2f (%s)\n\n"   + 
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
		uv.SunInfo.SunTimes.Sunrise,
		uv.SunInfo.SunTimes.SolarNoon,
		uv.SunInfo.SunTimes.Sunset,
		uv.SunInfo.SunTimes.Night, 
		uv.SunInfo.SunTimes.GoldenHour,
		uv.SunInfo.SunTimes.GoldenHourEnd,
		uv.SafeExposureTime.ST1, uv.SafeExposureTime.ST4, 
		uv.SafeExposureTime.ST2, uv.SafeExposureTime.ST5, 
		uv.SafeExposureTime.ST3, uv.SafeExposureTime.ST6,
	)
}
