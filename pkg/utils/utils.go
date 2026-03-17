// Package utils provides small shared helpers for HTTP handlers.
package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody reads the entire request body and unmarshals JSON into dest.
// Callers should pass a pointer to the target struct.
func ParseBody(r *http.Request, dest interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.Unmarshal(body, dest)
}
