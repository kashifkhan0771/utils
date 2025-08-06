/*
Package fake provides utilities for generating fake data.
*/
package fake

import (
	cryptrnd "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"time"

	rnd "github.com/kashifkhan0771/utils/rand"
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
	_, err := io.ReadFull(cryptrnd.Reader, uuid)
	if err != nil {
		return "", fmt.Errorf("failed to generate random UUID: %w", err)
	}

	// Set the version to 4
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	// Set the variant to 2 (RFC 4122)
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	uuidStr := hex.EncodeToString(uuid)

	return uuidStr[0:8] + "-" + uuidStr[8:12] + "-" + uuidStr[12:16] + "-" + uuidStr[16:20] + "-" + uuidStr[20:32], nil
}

// RandomDate generates a random date.
func RandomDate() (time.Time, error) {
	start := time.Date(EpochYear, time.Month(EpochMonth), EpochDay, EpochHour, EpochMinute, EpochSecond, EpochNano, time.UTC).UnixNano()
	end := time.Now().UnixNano()

	nanos, err := rnd.NumberInRange(start, end)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to generate random time: %w", err)
	}

	return time.Unix(0, nanos), nil
}

// RandomPhoneNumber generates a random phone number.
func RandomPhoneNumber() (string, error) {
	areaCode, err := rnd.NumberInRange(100, 999)
	if err != nil {
		return "", fmt.Errorf("failed to generate areaCode: %w", err)
	}

	firstPart, err := rnd.NumberInRange(100, 999)
	if err != nil {
		return "", fmt.Errorf("failed to generate firstPart: %w", err)
	}

	secondPart, err := rnd.NumberInRange(1000, 9999)
	if err != nil {
		return "", fmt.Errorf("failed to generate secondPart: %w", err)
	}

	return fmt.Sprintf("+1 (%03d) %03d-%04d", areaCode, firstPart, secondPart), nil
}

// RandomAddress generates a random address.
func RandomAddress() (string, error) {
	streetNumber, err := rnd.NumberInRange(1, MaxStreetNumber)
	if err != nil {
		return "", fmt.Errorf("failed to generate street number: %w", err)
	}

	streetName, err := rnd.Pick(streetNames)
	if err != nil {
		return "", fmt.Errorf("failed to pick street name: %w", err)
	}

	city, err := rnd.Pick(cities)
	if err != nil {
		return "", fmt.Errorf("failed to pick city: %w", err)
	}

	state, err := rnd.Pick(states)
	if err != nil {
		return "", fmt.Errorf("failed to pick state: %w", err)
	}

	postalCode, err := rnd.Pick(postalCodes)
	if err != nil {
		return "", fmt.Errorf("failed to pick postal code: %w", err)
	}

	return fmt.Sprintf("%d %s, %s, %s %s, USA", streetNumber, streetName, city, state, postalCode), nil
}
