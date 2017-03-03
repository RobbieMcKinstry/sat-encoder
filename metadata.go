package encoder

func NewMetadata(owner interface{}) *Metadata {
	return &Metadata{
		owner: owner,
		data:  make(map[string]interface{}),
	}
}

func (m *Metadata) Satisfies(p Proposition) bool {
	return p(m)
}

func (m *Metadata) AddLabel(label string, val interface{}) error {
	if _, ok := m.data[label]; ok {
		return LabelExistsError{label}
	}
	m.data[label] = val
	return nil
}

func (m *Metadata) LabelMatches(label string, val interface{}) bool {
	if v, ok := m.data[label]; ok {
		return v == val
	}
	return false
}

func (m *Metadata) LabelMatchesWith(label string, val interface{}, matcher func(v1, v2 interface{}) bool) bool {
	if v, ok := m.data[label]; ok {
		return matcher(v, val)
	}
	return false
}

func (m *Metadata) LabelMatchesString(label, expected string) bool {
	if val, ok := m.data[label]; ok {
		if observed, ok := IsString(val); ok {
			return observed == expected
		}
	}
	return false
}

func (m *Metadata) LabelEqualsInt(label string, expected int) bool {
	if val, ok := m.data[label]; ok {
		if observed, ok := IsInt(val); ok {
			return observed == expected
		}
	}
	return false
}

func (m *Metadata) LabelLessThanInt(label string, expected int) bool {
	if val, ok := m.data[label]; ok {
		if observed, ok := IsInt(val); ok {
			return observed < expected
		}
	}
	return false
}

func (m *Metadata) LabelGreaterThanInt(label string, expected int) bool {
	if val, ok := m.data[label]; ok {
		if observed, ok := IsInt(val); ok {
			return observed > expected
		}
	}
	return false
}

func (m *Metadata) LabelGEqInt(label string, expected int) bool {
	if val, ok := m.data[label]; ok {
		if observed, ok := IsInt(val); ok {
			return observed >= expected
		}
	}
	return false
}

func (m *Metadata) LabelLEqInt(label string, expected int) bool {
	if val, ok := m.data[label]; ok {
		if observed, ok := IsInt(val); ok {
			return observed <= expected
		}
	}
	return false
}

func (m *Metadata) LabelNEqInt(label string, expected int) bool {
	if val, ok := m.data[label]; ok {
		if observed, ok := IsInt(val); ok {
			return observed != expected
		}
	}
	return false
}

func IsString(val interface{}) (string, bool) {
	switch val := val.(type) {
	case string:
		return val, true
	}
	return "", false
}

func IsInt(val interface{}) (int, bool) {
	switch val := val.(type) {
	case int:
		return val, true
	}
	return 0, false
}
