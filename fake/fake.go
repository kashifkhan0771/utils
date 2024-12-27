/*
Package fake provides utilities for generating fake data.
*/
package fake

import (
	"crypto/rand"
	"fmt"
	rnd "github.com/kashifkhan0771/utils/rand"
	"io"
	"time"
)

const (
	EpochYear   = 1970
	EpochMonth  = 1
	EpochDay    = 1
	EpochHour   = 0
	EpochMinute = 0
	EpochSecond = 0
	EpochNano   = 0
)

const MaxStreetNumber = 150

var streetNames = []string{"Main St", "High St", "Broadway", "Maple Ave", "Oak St", "Pine St", "Cedar St", "Elm St"}
var cities = []string{"Springfield", "Rivertown", "Lakeview", "Greenville", "Fairview", "Madison", "Georgetown", "Clinton"}
var states = []string{"AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY"}
var postalCodes = []string{"12345", "23456", "34567", "45678", "56789", "67890", "78901", "89012", "90123"}

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

// RandomDate generates a random date.
func RandomDate() (time.Time, error) {
	start := time.Date(EpochYear, time.Month(EpochMonth), EpochDay, EpochHour, EpochMinute, EpochSecond, EpochNano, time.UTC).Unix()
	end := time.Now().Unix()

	sec, err := rnd.NumberInRange(start, end)
	if err != nil {
		return time.Time{}, err
	}

	nsec, err := rnd.NumberInRange(0, 1e9)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(sec, nsec), nil
}

// RandomPhoneNumber generates a random phone number.
func RandomPhoneNumber() (string, error) {
	areaCode, err := rnd.NumberInRange(100, 999)
	if err != nil {
		return "", err
	}

	firstPart, err := rnd.NumberInRange(100, 999)
	if err != nil {
		return "", err
	}

	secondPart, err := rnd.NumberInRange(1000, 9999)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("+1 (%d) %d-%d", areaCode, firstPart, secondPart), nil
}

// RandomAddress generates a random address.
func RandomAddress() (string, error) {
	streetNumber, err := rnd.NumberInRange(0, MaxStreetNumber)
	if err != nil {
		return "", err
	}

	idx, err := rnd.NumberInRange(0, int64(len(streetNames)-1))
	if err != nil {
		return "", err
	}
	streetName := streetNames[idx]

	idx, err = rnd.NumberInRange(0, int64(len(streetNames)-1))
	if err != nil {
		return "", err
	}
	city := cities[idx]

	idx, err = rnd.NumberInRange(0, int64(len(states)-1))
	if err != nil {
		return "", err
	}
	state := states[idx]

	idx, err = rnd.NumberInRange(0, int64(len(postalCodes)-1))
    if err != nil {
        return "", err
    }
    postalCode := postalCodes[idx]

    return fmt.Sprintf("%d %s, %s, %s %s, USA", streetNumber, streetName, city, state, postalCode), nil
}
