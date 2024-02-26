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
			state := NewStateMap()
			state.SetState("key1", true)

			if got := state.HasState(tt.args.stateType); got != tt.want {
				t.Errorf("HasState() = %v, want %v", got, tt.want)
			}
		})
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
			m:    map[string]string{"key1": "value2"},
			args: args{key: "key1", value: "value1"},
		},
		{
			name: "success - add a new key and value",
			m:    map[string]string{"key1": "value2"},
			args: args{key: "key2", value: "value1"},
		},
		{
			name: "success - update key and value with nil map",
			m:    nil,
			args: args{key: "key1", value: "value1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Update(tt.args.key, tt.args.value)
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
			m:    map[string]string{"key1": "value1"},
			args: args{key: "key1"},
			want: true,
		},
		{
			name: "success - key does not exist",
			m:    map[string]string{"key1": "value1"},
			args: args{key: "key2"},
			want: false,
		},
		{
			name: "success - check in nil map",
			m:    nil,
			args: args{key: "key2"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Has(tt.args.key); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
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
			m:    map[string]string{"key1": "value1"},
			args: args{key: "key1"},
			want: "value1",
		},
		{
			name: "success - get a value of the non existing key",
			m:    map[string]string{"key1": "value1"},
			args: args{key: "key2"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Value(tt.args.key); got != tt.want {
				t.Errorf("Value() = %v, want %v", got, tt.want)
			}
		})
	}
}
