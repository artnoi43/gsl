package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// GetAndParse fetches data from the HTTP endpoint,
// and parses response's JSON body into the interface.
func GetAndParse(u string, v interface{}) error {
	resp, err := http.Get(u)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"failed to get url: %s",
				u,
			),
		)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusTooManyRequests:
		return ErrRateLimitExceeded
	default:
		return errors.New("non-200 status code")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"failed to read body: %s",
				resp.Body,
			),
		)
	}
	return json.Unmarshal(body, &v)
}
