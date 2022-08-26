package http

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// GetAndParse fetches data from the HTTP endpoint,
// and parses response's JSON body into the interface.
func GetAndParse(u string, v interface{}) error {
	client := http.Client{
		Timeout: time.Second * 3,
	}
	resp, err := client.Get(u)
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

	body, err := io.ReadAll(resp.Body)
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
