package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no auth header included in request")

func GetApiKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splithAuth := strings.Split(authHeader, " ")
	if len(splithAuth) < 2 || splithAuth[0] != "ApiKey" {
		return "", errors.New("malformed authoriztion header")
	}

	return splithAuth[1], nil
}
