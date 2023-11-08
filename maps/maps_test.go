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
