package main

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
)

const (
	DEFAULT_SUN_TIMES_FORMAT string = "15:04"
	HEADER_TIME_FORMAT       string = "2006-02-01 15:04 UTC"
	OPEN_UV_TIME_FORMAT      string = "2006-02-01T15:04:05.999Z"
)

/* Ensures the path to resource directory is correct when running from PATH (different WD). 
 * Accepts fname being the relative path to the file from the app's root (executable location).
 */
 func GetResPath(fname string) (string, error) {
	exePath, err := exec.LookPath(filepath.Base(os.Args[0]))
	if err != nil {
		return "", err
	}
	
	rootDir, _ := filepath.Split(exePath)
	return filepath.FromSlash(path.Join(rootDir[:len(rootDir) - 1], fname)), nil
}

/* Reparses the timestamp from format used by OpenUV API into newFormat. */
func ReformatTime(timestamp, newFormat string) string {
	t, _ := time.Parse(OPEN_UV_TIME_FORMAT, timestamp)
	return t.Format(newFormat)
}
