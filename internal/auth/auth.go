package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns the API key from the request headers.
// If the API key is not found, an error is returned.
// If the API key is found, it is returned as a string.
// If the API key is found, but is not valid, an error is returned.
// Example:
// Authorization: ApiKey {insert API key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid Authorization header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of Authorization header")
	}
	if vals[1] == "" {
		return "", errors.New("missing API key")
	}
	return vals[1], nil
}
