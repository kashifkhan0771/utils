package fake

import (
	"regexp"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
    uuid, err := RandomUUID()
    if err != nil {
        t.Fatalf("Expected no error, got %v", err)
    }

    // Test if UUID is of correct length
    if len(uuid) != 36 {
        t.Errorf("Expected length 36, got %d", len(uuid))
    }

    // Test if UUID matches the correct format
    isValidUUID := regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`).MatchString
    if !isValidUUID(uuid) {
        t.Errorf("UUID %s does not match the required format", uuid)
    }

    // Test for uniqueness
    uuidSet := make(map[string]struct{})
    for i := 0; i < 1000; i++ {
        uuid, err := RandomUUID()
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }
        if _, exists := uuidSet[uuid]; exists {
            t.Errorf("Duplicate UUID found: %s", uuid)
        }
        uuidSet[uuid] = struct{}{}
    }
}