package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// GetAndParse fetches data from the HTTP endpoint,
// and parses response's JSON body into the interface.
func GetAndParse(ctx context.Context, url string, v interface{}) error {
	client := http.Client{
		Timeout: time.Second * 3,
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return errors.Wrapf(err, "failed to build request for url %s", url)
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(
			err,
			fmt.Sprintf(
				"failed to get url: %s",
				url,
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
	return errors.Wrapf(json.Unmarshal(body, &v), "failed to unmarshal body from %s", url)
}
