package main

import (
	"encoding/binary"
	"errors"
	"os"
	"path"
	"time"
)

const (
	CACHE_DIR     string = "cache"

	// Path to request counter's cache file
	CACHE_RC      string = CACHE_DIR + "/uv_rc"

	// Expected length of the request counter's cache file
	EXPECT_RC_LEN int    = 16
)

var (
	// Error returned when cache file length is different than expected
	errBadCache   error = errors.New("cache file corrupted")
)

/* Request counter for OpenUV API. It stores the information about the number of requests 
 * left until the end of the day. The API resets the requests quota at midnight UTC.
 */
type RequestCounter struct {
	// Requests left until reset
	Remaining int64

	// Expiration time, after which Remainig is set to settings.RequestLimit. 
	Expires   int64
}

/* Returns true if rc has expired. */
func (rc *RequestCounter) Expired() bool {
	return time.Now().UTC().Unix() > rc.Expires
}

/* Sets the new expiration time to next midnight UTC. */
func (rc *RequestCounter) SetExpirationTime() {
	now := time.Now().UTC()
	et  := time.Date(now.Year(), now.Month(), now.Day() + 1, 0, 0, 0, 0, now.Location())

	rc.Expires = et.Unix()
}

/* Returns the remaining number of requests stored within cache file. */
func ProcessRequestCounter(s *Settings) (int64, error) {
	rc, err := ReadRequestCounter()
	if err != nil || rc.Expired() {
		rc.Remaining = s.RequestLimit
		rc.SetExpirationTime()
	}

	if rc.Remaining == 0 {
		return rc.Remaining, nil
	}

	rc.Remaining--

	err = SaveRequestCounter(rc)
	if err != nil {
		return rc.Remaining, err
	}

	return rc.Remaining, nil
}

/* Retrieves RequestCounter from cache file. */
func ReadRequestCounter() (*RequestCounter, error) {
	var rc RequestCounter

	stream, err := os.ReadFile(path.Join(ROOT_DIR, CACHE_RC))
	if err != nil {
		return &rc, err
	} else if len(stream) < EXPECT_RC_LEN {
		return &rc, errBadCache
	}

	rc.Expires   = int64(binary.BigEndian.Uint64(stream[:8]))
	rc.Remaining = int64(binary.BigEndian.Uint64(stream[8:]))

	return &rc, nil
}

/* Saves RequestCounter to cache file. */
func SaveRequestCounter(rc *RequestCounter) error {
	if _, err := os.Stat(path.Join(ROOT_DIR, CACHE_DIR)); os.IsNotExist(err) {
		err := os.Mkdir(path.Join(ROOT_DIR, CACHE_DIR), 0775)
		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path.Join(ROOT_DIR, CACHE_RC), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0775)
	if err != nil {
		return err
	}
	defer f.Close()

	b := make([]byte, 16)

	binary.BigEndian.PutUint64(b[:8], uint64(rc.Expires))
	binary.BigEndian.PutUint64(b[8:], uint64(rc.Remaining))

	f.Write(b)

	return nil
}
