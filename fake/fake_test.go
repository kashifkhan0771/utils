package fake

import (
	"regexp"
	"testing"
	"testing/quick"
	"time"
)

func TestRandomUUID(t *testing.T) {
	t.Parallel()

	re := regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`)
	seen := make(map[string]struct{})

	prop := func() bool {
		u, err := RandomUUID()
		if err != nil || len(u) != 36 || !re.MatchString(u) {
			return false
		}
		// simple uniqueness check across trials
		if _, dup := seen[u]; dup {
			return false
		}
		seen[u] = struct{}{}
		return true
	}

	if err := quick.Check(prop, &quick.Config{MaxCount: 1_000}); err != nil {
		t.Fatalf("RandomUUID property failed: %v", err)
	}
}

func TestRandomDate(t *testing.T) {
	t.Parallel()

	start := time.Date(EpochYear, time.Month(EpochMonth), EpochDay, EpochHour, EpochMinute, EpochSecond, EpochNano, time.UTC)

	prop := func() bool {
		d, err := RandomDate()
		if err != nil {
			return false
		}
		end := time.Now()
		return !d.Before(start) && !d.After(end)
	}

	if err := quick.Check(prop, &quick.Config{MaxCount: 1_000}); err != nil {
		t.Fatalf("RandomDate property failed: %v", err)
	}
}

// Sanity property: RandomDate never returns the zero time.
func TestRandomDate_NonZero_Property(t *testing.T) {
	t.Parallel()

	prop := func() bool {
		d, err := RandomDate()
		return err == nil && !d.IsZero()
	}

	if err := quick.Check(prop, &quick.Config{MaxCount: 1_000}); err != nil {
		t.Fatalf("RandomDate_NonZero property failed: %v", err)
	}
}

func TestRandomPhoneNumber(t *testing.T) {
	t.Parallel()

	re := regexp.MustCompile(`^\+1 \(\d{3}\) \d{3}-\d{4}$`)

	prop := func() bool {
		phoneNumber, err := RandomPhoneNumber()

		return err == nil && re.MatchString(phoneNumber)
	}

	if err := quick.Check(prop, &quick.Config{MaxCount: 1_000}); err != nil {
		t.Fatalf("RandomPhoneNumber property failed: %v", err)
	}
}

func TestRandomAddress(t *testing.T) {
	t.Parallel()

	re := regexp.MustCompile(`^\d+ [A-Za-z ]+, [A-Za-z ]+, [A-Z]{2} \d{5}, USA$`)

	prop := func() bool {
		address, err := RandomAddress()

		return err == nil && re.MatchString(address)
	}

	if err := quick.Check(prop, &quick.Config{MaxCount: 1_000}); err != nil {
		t.Fatalf("RandomAddress property failed: %v", err)
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkGenerateUUID(b *testing.B) {
	for b.Loop() {
		_, _ = RandomUUID()
	}
}

func BenchmarkRandomDate(b *testing.B) {
	for b.Loop() {
		_, _ = RandomDate()
	}
}

func BenchmarkRandomPhoneNumber(b *testing.B) {
	for b.Loop() {
		_, _ = RandomPhoneNumber()
	}
}

func BenchmarkRandomAddress(b *testing.B) {
	for b.Loop() {
		_, _ = RandomAddress()
	}
}
