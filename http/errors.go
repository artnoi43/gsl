package http

import "github.com/pkg/errors"

var ErrRateLimitExceeded = errors.New("rate limit exceeded")
