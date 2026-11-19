package ctxutils

import (
	"context"
	"testing"
)

func TestStringValue(t *testing.T) {
	tests := []struct {
		name      string
		setKey    ContextKeyString
		getKey    ContextKeyString
		value     string
		wantValue string
		wantOK    bool
	}{
		{
			name:      "same key returns stored value",
			setKey:    ContextKeyString{Key: "userID"},
			getKey:    ContextKeyString{Key: "userID"},
			value:     "Shahzad",
			wantValue: "Shahzad",
			wantOK:    true,
		},
		{
			name:      "wrong key is not found",
			setKey:    ContextKeyString{Key: "userID"},
			getKey:    ContextKeyString{Key: "wrongKey"},
			value:     "Shahzad",
			wantValue: "",
			wantOK:    false,
		},
		{
			name:      "empty string is stored and found",
			setKey:    ContextKeyString{Key: "userID"},
			getKey:    ContextKeyString{Key: "userID"},
			value:     "",
			wantValue: "",
			wantOK:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := SetStringValue(context.Background(), tt.setKey, tt.value)

			got, ok := GetStringValue(ctx, tt.getKey)
			if ok != tt.wantOK {
				t.Errorf("GetStringValue() found = %v, want %v", ok, tt.wantOK)
			}
			if got != tt.wantValue {
				t.Errorf("GetStringValue() = %q, want %q", got, tt.wantValue)
			}
		})
	}
}

func TestIntValue(t *testing.T) {
	tests := []struct {
		name      string
		setKey    ContextKeyInt
		getKey    ContextKeyInt
		value     int
		wantValue int
		wantOK    bool
	}{
		{
			name:      "same key returns stored value",
			setKey:    ContextKeyInt{Key: 123},
			getKey:    ContextKeyInt{Key: 123},
			value:     456,
			wantValue: 456,
			wantOK:    true,
		},
		{
			name:      "wrong key is not found",
			setKey:    ContextKeyInt{Key: 123},
			getKey:    ContextKeyInt{Key: 999},
			value:     456,
			wantValue: 0,
			wantOK:    false,
		},
		{
			name:      "zero is stored and found",
			setKey:    ContextKeyInt{Key: 123},
			getKey:    ContextKeyInt{Key: 123},
			value:     0,
			wantValue: 0,
			wantOK:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx := SetIntValue(context.Background(), tt.setKey, tt.value)

			got, ok := GetIntValue(ctx, tt.getKey)
			if ok != tt.wantOK {
				t.Errorf("GetIntValue() found = %v, want %v", ok, tt.wantOK)
			}
			if got != tt.wantValue {
				t.Errorf("GetIntValue() = %d, want %d", got, tt.wantValue)
			}
		})
	}
}

func TestMissingValue(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	if _, ok := GetStringValue(ctx, ContextKeyString{Key: "userID"}); ok {
		t.Errorf("GetStringValue() on empty context found a value, want not found")
	}

	if _, ok := GetIntValue(ctx, ContextKeyInt{Key: 123}); ok {
		t.Errorf("GetIntValue() on empty context found a value, want not found")
	}
}

// ================================================================================
// ### BENCHMARKS
// ================================================================================

func BenchmarkSettingAndGettingStringKey(b *testing.B) {
	var (
		ctx   context.Context
		key   = ContextKeyString{Key: "id"}
		value = "value-123"
	)

	b.ReportAllocs()
	for b.Loop() {
		ctx = context.Background()
		ctx = SetStringValue(ctx, key, value)
		_, _ = GetStringValue(ctx, key)
	}
}

func BenchmarkSettingAndGettingIntKey(b *testing.B) {
	var (
		ctx   context.Context
		key   = ContextKeyInt{Key: 0}
		value = 123
	)

	b.ReportAllocs()
	for b.Loop() {
		ctx = context.Background()
		ctx = SetIntValue(ctx, key, value)
		_, _ = GetIntValue(ctx, key)
	}
}
