package main

import (
	"os"
	"path/filepath"
	"time"
)

const (
	DEFAULT_SUN_TIMES_FORMAT string = "15:04"
	HEADER_TIME_FORMAT       string = "2006-01-02 15:04 UTC"
	OPEN_UV_TIME_FORMAT      string = "2006-01-02T15:04:05.999Z"
)

/* Returns the root directory in which the executable is located. */
func GetRootDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(exePath), nil
}

/* Converts timestamp from format used by OpenUV API into newFormat. */
func ReformatTime(timestamp, newFormat string) string {
	t, _ := time.Parse(OPEN_UV_TIME_FORMAT, timestamp)
	return t.Format(newFormat)
}
