package main

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
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
