package maps

import "testing"

func TestStateMap_IsState(t *testing.T) {
	type args struct {
		stateType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - found a state in state map",
			args: args{stateType: "key1"},
			want: true,
		},
		{
			name: "fail - not found a state in state map",
			args: args{stateType: "key2"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			state := NewStateMap()
			state.SetState("key1", true)

			if got := state.IsState(tt.args.stateType); got != tt.want {
				t.Errorf("IsState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStateMap_HasState(t *testing.T) {
	type args struct {
		stateType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - found the state in the state map",
			args: args{stateType: "key1"},
			want: true,
		},
		{
			name: "fail - not found the state in the state map",
			args: args{stateType: "key2"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			state := NewStateMap()
			state.SetState("key1", true)

			if got := state.HasState(tt.args.stateType); got != tt.want {
				t.Errorf("HasState() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStateMap_IsState_vs_HasState_FalseValue(t *testing.T) {
	state := NewStateMap()
	state.SetState("keyFalse", false)

	if got := state.IsState("keyFalse"); got {
		t.Fatalf("IsState(keyFalse) = %v, want false", got)
	}
	if got := state.HasState("keyFalse"); !got {
		t.Fatalf("HasState(keyFalse) = %v, want true", got)
	}
}

func TestStateMap_ToggleState(t *testing.T) {
	type args struct {
		stateType string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success - toggled the state in the state map",
			args: args{stateType: "key1"},
			want: false,
		},
		{
			name: "success - toggled the state in the state map",
			args: args{stateType: "key2"},
			want: true,
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			state := NewStateMap()
			state.SetState(tt.args.stateType, !tt.want)

			state.ToggleState(tt.args.stateType)

			got := state.IsState(tt.args.stateType)
			if got != tt.want {
				t.Errorf("ToggleState() : %v, want : %v", got, tt.want)
			}
		})
	}
}

func TestStateMap_ToggleState_TwiceReturnsOriginal(t *testing.T) {
	state := NewStateMap()
	state.SetState("k", true)

	state.ToggleState("k") // true -> false
	state.ToggleState("k") // false -> true

	if got := state.IsState("k"); !got {
		t.Fatalf("want state true after two toggles, got false")
	}
}

func TestMetadata_Update(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		m    Metadata
		args args
	}{
		{
			name: "success - update key and value",
			m:    Metadata{map[string]string{"key1": "value2"}},
			args: args{key: "key1", value: "value1"},
		},
		{
			name: "success - add a new key and value",
			m:    Metadata{map[string]string{"key1": "value2"}},
			args: args{key: "key2", value: "value1"},
		},
		{
			name: "success - update key and value with nil map",
			m:    Metadata{},
			args: args{key: "key1", value: "value1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Get value before updating
			before := tt.m.Value(tt.args.key)

			// Update value
			tt.m.Update(tt.args.key, tt.args.value)

			// Get value after updating
			after := tt.m.Value(tt.args.key)

			// Check if the value was updated correctly
			if after != tt.args.value || (before == tt.args.value && after == tt.args.value) {
				t.Errorf("Update() = %v, want %v", after, tt.args.value)
			}
		})
	}
}

func TestMetadata_Has(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    Metadata
		args args
		want bool
	}{
		{
			name: "success - key exist",
			m:    Metadata{map[string]string{"key1": "value1"}},
			args: args{key: "key1"},
			want: true,
		},
		{
			name: "success - key does not exist",
			m:    Metadata{map[string]string{"key1": "value1"}},
			args: args{key: "key2"},
			want: false,
		},
		{
			name: "success - check in nil map",
			m:    Metadata{nil},
			args: args{key: "key2"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.Has(tt.args.key); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMetadata_Has_AfterUpdateFromNil(t *testing.T) {
	var m Metadata // nil map inside
	m.Update("k", "v")

	if !m.Has("k") {
		t.Fatalf("Has(k) = false after Update on nil map; want true")
	}
	if got := m.Value("k"); got != "v" {
		t.Fatalf("Value(k) = %q, want %q", got, "v")
	}
}

func TestMetadata_Value(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		m    Metadata
		args args
		want string
	}{
		{
			name: "success - get a value of the key",
			m:    Metadata{map[string]string{"key1": "value1"}},
			args: args{key: "key1"},
			want: "value1",
		},
		{
			name: "success - get a value of the non existing key",
			m:    Metadata{map[string]string{"key1": "value1"}},
			args: args{key: "key2"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.m.Value(tt.args.key); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
