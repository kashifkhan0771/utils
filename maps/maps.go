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
	return s[stateType]
}

// HasState check if particular state is present in the StateMap
func (s StateMap) HasState(stateType string) bool {
	_, ok := s[stateType]

	return ok
}

type metadata map[string]string

type Metadata struct {
	metadata
}

func NewMetadata() Metadata {
	return Metadata{make(map[string]string)}
}

func (m *Metadata) Update(key, value string) {
	ensureMetadata(m)

	m.metadata[key] = value
}

func (m *Metadata) Has(key string) bool {
	ensureMetadata(m)

	_, ok := m.metadata[key]

	return ok
}

func (m *Metadata) Value(key string) string {
	ensureMetadata(m)

	return m.metadata[key]
}

func ensureMetadata(m *Metadata) {
	if m.metadata == nil {
		m.metadata = make(map[string]string)
	}
}
