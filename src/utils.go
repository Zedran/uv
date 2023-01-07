package main

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	DEFAULT_SUN_TIMES_FORMAT string = "15:04"
	HEADER_TIME_FORMAT       string = "2006-01-02 15:04 UTC"
	OPEN_UV_TIME_FORMAT      string = "2006-01-02T15:04:05.999Z"
)

var errCoordinateOverflow    error  = errors.New("coordinate value out of range")

/* Converts geographical coordinate from string to float. Errors are returned
 * in case of parsing problems or if the coordinate's value limit was exceeded.
 * Value limits should correspond to the coordinate type (lat == 90, lon == 180).
 */
func ConvertCoordinate(coordString string, limit float64) (float64, error) {
	coordNum, err := strconv.ParseFloat(coordString, 64)
	if err != nil {
		return 0, err
	}

	if math.Abs(coordNum) > limit {
		return 0, errCoordinateOverflow
	}

	return coordNum, nil
}

/* Returns the root directory in which the executable is located. */
func GetRootDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(exePath), nil
}

/* Converts degrees to radians. */
func Rad(deg float64) float64 {
	return deg * math.Pi / 180
}

/* Converts timestamp from format used by OpenUV API into newFormat. */
func ReformatTime(timestamp, newFormat string) string {
	t, _ := time.Parse(OPEN_UV_TIME_FORMAT, timestamp)
	return t.Format(newFormat)
}
