package ctxutils

import (
	"context"
	"fmt"
	"testing"
)

// TODO: these can be written as a table test

func TestSetStringValueAndGetStringValue(t *testing.T) {
	ctx := context.Background()

	key := ContextKeyString{Key: "userID"}
	value := "Shahzad"

	ctx = SetStringValue(ctx, key, value)
	retrievedValue, ok := GetStringValue(ctx, key)
	if !ok {
		t.Errorf("Expected value to be found, but it was not.")
	}

	if retrievedValue != value {
		t.Errorf("Expected value: %s, but got: %v", value, retrievedValue)
	}
}

func TestSetIntValueAndGetIntValue(t *testing.T) {
	ctx := context.Background()

	key := ContextKeyInt{Key: 123}
	value := 456

	ctx = SetIntValue(ctx, key, value)
	retrievedValue, ok := GetIntValue(ctx, key)
	if !ok {
		t.Errorf("Expected value to be found, but it was not.")
	}

	if retrievedValue != value {
		t.Errorf("Expected value: %d, but got: %v", value, retrievedValue)
	}
}

func TestSetStringValueWithWrongKey(t *testing.T) {
	ctx := context.Background()

	key := ContextKeyString{Key: "userID"}
	value := "Shahzad"
	ctx = SetStringValue(ctx, key, value)

	wrongKey := ContextKeyString{Key: "wrongKey"}
	_, ok := GetStringValue(ctx, wrongKey)
	if ok {
		t.Errorf("Expected value not to be found, but it was.")
	}
}

func TestSetIntValueWithWrongKey(t *testing.T) {
	ctx := context.Background()

	key := ContextKeyInt{Key: 123}
	value := 456

	ctx = SetIntValue(ctx, key, value)
	wrongKey := ContextKeyInt{Key: 999}
	_, ok := GetIntValue(ctx, wrongKey)

	if ok {
		t.Errorf("Expected value not to be found, but it was.")
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkSettingAndGettingStringKey(b *testing.B) {
	ctx := context.Background()
	key := ContextKeyString{Key: "id"}

	b.ReportAllocs()
	for i := 0; b.Loop(); i++ {
		ctx = SetStringValue(ctx, key, fmt.Sprintf("value-%d", i))
		_, _ = GetStringValue(ctx, key)
	}
}

func BenchmarkSettingAndGettingIntKey(b *testing.B) {
	ctx := context.Background()
	key := ContextKeyInt{Key: 0}

	b.ReportAllocs()
	for i := 0; b.Loop(); i++ {
		ctx = SetIntValue(ctx, key, i)
		_, _ = GetIntValue(ctx, key)
	}
}
