package maps

// StateMap hold key as string and value in bool
type StateMap map[string]bool

// NewStateMap create new state.
func NewStateMap() StateMap {
	return make(StateMap)
}

// SetState sets a new state with provided value in the StateMap
func (s StateMap) SetState(stateType string, value bool) {
	s[stateType] = value
}

// IsState return the value of the state provided from StateMap
func (s StateMap) IsState(stateType string) bool {
	if ok, v := s[stateType]; ok {
		return v
	}

	return false
}

// HasState check if particular state is present in the StateMap
func (s StateMap) HasState(stateType string) bool {
	ok, _ := s[stateType]

	return ok
}
