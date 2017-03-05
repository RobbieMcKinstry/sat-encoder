package encoder

import (
	"github.com/mitchellh/go-sat/cnf"
)

type (
	Metadata struct {
		data map[string]interface{}
		// owner is the object that this metadata references
		owner interface{}
		name  string
	}

	Constrainer interface {
		AddLabel(string, interface{}) error
		Satisfies(Constraint) bool
		Encode([]Constraint) cnf.Clause
		GetMetadata() *Metadata
		Name() string
	}

	LabelExistsError struct {
		label string
	}

	Proposition func(*Metadata) bool

	Constraint struct {
		uid  int
		prop Proposition
	}

	setElem struct {
		lit  cnf.Lit
		meta *Metadata
		name string
	}

	stringTable map[cnf.Lit]*setElem

	solutionElem struct {
		c        Constrainer
		selected bool
	}
)

// For a constraint C,
// Find the Propositions that do satisfy C
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
var UID func() int = makeUIDGenerator()
