package web

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	versionFile = "version.txt"
)

// GetVersion gets the current version of the app for output in various places
// In order that this works without the whole git environment, this function
// looks for the version in a file called 'version.txt'

func GetVersion() string {
	b, err := ioutil.ReadFile(versionFile)
	if err != nil {
		return "not available"
	}

	return strings.TrimSpace(string(b))
}

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	RespondString(w, http.StatusOK, GetVersion())
}
