package utils

import (
	"bytes"
	"crypto/sha1" //nolint:gosec // sha1 is not cryptographically secure, but that is not a goal. we require predictable and uniform text transformation.
	"io"
	"net/url"
)

const (
	maxURLKeyLength = 200
)

// URLEscape escapes url with prefix
func URLEscape(prefix, u string) string {
	key := url.QueryEscape(u)
	if len(key) > maxURLKeyLength {
		h := sha1.New() //nolint:gosec // sha1 is not cryptographically secure, but that is not a goal.
		_, _ = io.WriteString(h, u)
		key = string(h.Sum(nil))
	}

	var buffer bytes.Buffer
	_, _ = buffer.WriteString(prefix)
	_, _ = buffer.WriteString(":")
	_, _ = buffer.WriteString(key)

	return buffer.String()
}
