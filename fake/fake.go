/*
Package fake provides utilities for generating fake data.
*/
package fake

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
)

// RandomUUID generates a fake UUIDv4.
func RandomUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, uuid)
	if err != nil {
		return "", err
	}

	// Set the version to 4
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	// Set the variant to 2 (RFC 4122)
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x", uuid[:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func RandomDate() time.Time {
	return time.Time{}
}
