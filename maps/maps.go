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

// ToggleState toggles the value of the state with the provided key in the StateMap.
func (s StateMap) ToggleState(stateType string) {
	s[stateType] = !s[stateType]
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
	ok := s[stateType]

	return ok
}

type Metadata map[string]string

func NewMetadata() Metadata {
	return make(map[string]string)
}

func (m Metadata) Update(key, value string) {
	if m == nil {
		m = NewMetadata()
	}

	m[key] = value
}

func (m Metadata) Has(key string) bool {
	_, ok := m[key]

	return ok
}

func (m Metadata) Value(key string) string {
	if !m.Has(key) {
		return ""
	}

	return m[key]
}
