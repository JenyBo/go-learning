package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("API key is not found")
	}
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid API key")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("wrong API key type")
	}

	return vals[1], nil
}
