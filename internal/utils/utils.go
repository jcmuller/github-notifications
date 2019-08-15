package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	userNameVar = "GITHUB_NOTIFICATIONS_USER_NAME"
	passwordVar = "GITHUB_NOTIFICATIONS_PASSWORD"
)

func getVar(variable string) (value string) {
	var ok bool

	value, ok = os.LookupEnv(variable)
	if !ok {
		fmt.Fprintf(os.Stderr, "%s is required to run\n", variable)
		os.Exit(1)
	}

	return
}

// Request builds a nice http request
func Request(method string, body io.Reader) (r *http.Request, err error) {
	url := "https://api.github.com/notifications"

	r, err = http.NewRequest(method, url, body)
	if err != nil {
		return
	}

	user := getVar(userNameVar)
	password := getVar(passwordVar)

	r.SetBasicAuth(user, password)
	r.Header.Set("Content-Type", "application/json")

	return
}
