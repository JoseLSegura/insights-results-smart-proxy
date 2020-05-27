package server

import (
	"fmt"
	"regexp"
)

const (
	// MainEndpoint defines suffix of the root endpoint
	MainEndpoint = ""
)

// MakeURLToEndpoint creates URL to endpoint, use constants from file endpoints.go
func MakeURLToEndpoint(apiPrefix, endpoint string, args ...interface{}) string {
	re := regexp.MustCompile(`\{[a-zA-Z_0-9]+\}`)
	endpoint = re.ReplaceAllString(endpoint, "%v")
	return apiPrefix + fmt.Sprintf(endpoint, args...)
}
