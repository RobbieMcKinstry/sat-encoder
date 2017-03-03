package main

import "fmt"

type (
	Metadata map[string]interface{}

	Constrainer interface {
		Add(string, interface{}) error
		Satisfies(Proposition) bool
		Encode([]Proposition) []int
	}

	LabelExistsError struct {
		label string
	}

	Proposition func(Metadata) bool

	Constraint struct {
		uid             int
		truthVal        bool
		PropositionFunc Proposition
	}

	// func (Constraint) Encode() cnf.Clause {
	//     return []cnf.Lit{}
	// }
)

// For a constraint C,
// Find the Propositions that do not satisfy C
// Conjunct them with ¬C

// a -> arrays constraint
// l -> loop constraint
// x1 satisfies the loops constraint
// x2 satisfies the loops constraint
// x3 satisfies the arrays constraint
// x4 satisfies the arrays constraint

// Therefore, the formula is
//        a          ^
//        l          ^
// ( ¬ a ∨ x3 v x4 ) ^
// ( ¬ l v x1 ∨ x2 )

// I don't think you have to do this.
// Finally, encode each positive selection as a single clause, with each node
// followed by the inverse of what they satisfy:
// ( x1 ∨ ¬ l ∨ ¬ d )

func (m Metadata) EncodeProps(props []Proposition) map[Proposition]bool {
	var result = map[Proposition]bool{}
	for _, p := range props {
		result[p] = m.Satisfies(Proposition)
	}
	return result
}

func (m Metadata) Satisfies(p Proposition) {
	return p(m)
}

func StringProposition(label, val string) Proposition {
	return func(m Metadata) bool {
		return m.LabelMatchesString(label, val)
	}
}

func IntProposition(label, val int) Proposition {
	return func(m Metadata) bool {
		return m.LabelMatchesInt(label, val)
	}
}

func (l LabelExists) String() {
	return fmt.Sprintf("The label %s already exists.", l.label)
}

func (m Metadata) Add(label string, val interface{}) error {
	if _, ok := m[label]; ok {
		return LabelExists{label}
	}
	m[label] = val
	return nil
}

func (m Metadata) LabelMatches(label string, val interface{}) bool {
	if v, ok := m[label]; ok {
		return v == val
	}
	return false
}

func (m Metadata) LabelMatchesWith(label string, val interface{}, matcher func(v1, v2 interface{}) bool) bool {
	if v, ok := m[label]; ok {
		return matcher(v, val)
	}
	return false
}

func (m Metadata) LabelMatchesString(label, expected string) {
	if val, ok := m[label]; ok {
		if observed, ok := IsString(v); ok {
			return observed == expected
		}
	}
	return false
}

func (m Metadata) LabelEqualsInt(label, expected int) {
	if val, ok := m[label]; ok {
		if observed, ok := IsInt(v); ok {
			return observed == expected
		}
	}
	return false
}

func (m Metadata) LabelLessThanInt(label, expected int) {
	if val, ok := m[label]; ok {
		if observed, ok := IsInt(v); ok {
			return observed < expected
		}
	}
	return false
}

func (m Metadata) LabelGreaterThanInt(label, expected int) {
	if val, ok := m[label]; ok {
		if observed, ok := IsInt(v); ok {
			return observed > expected
		}
	}
	return false
}

func (m Metadata) LabelGEqInt(label, expected int) {
	if val, ok := m[label]; ok {
		if observed, ok := IsInt(v); ok {
			return observed >= expected
		}
	}
	return false
}

func (m Metadata) LabelLEqInt(label, expected int) {
	if val, ok := m[label]; ok {
		if observed, ok := IsInt(v); ok {
			return observed <= expected
		}
	}
	return false
}

func (m Metadata) LabelNEqInt(label, expected int) {
	if val, ok := m[label]; ok {
		if observed, ok := IsInt(v); ok {
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
	return nil, false
}

func IsInt(val interface{}) (int, bool) {
	switch val := val.(type) {
	case int:
		return val, true
	}
	return nil, false
}
