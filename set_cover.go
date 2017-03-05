package encoder

import (
	"github.com/mitchellh/go-sat"
	"github.com/mitchellh/go-sat/cnf"
)

// TODO make a global string table mapping literals to their names.
// TODO Refact the SetCover method to be a lot smaller
func SetCover(dataSet []Constrainer, constraints []Constraint) ([]Constrainer, bool) {

	var (
		// Convert the constrainter into a simplier datatype
		litSet []*setElem = toLitSet(dataSet)

		// the table holding the mapping from cnf.Lit -> setElem
		st = make(stringTable)

		// add each constraint's individual identifier as it's own clause
		clauses = buildUnitConstraints(constraints, st)
	)
	// add each of the literals to the string table
	st.addLiterals(litSet)

	// build the constraint clauses.
	constraintClauses := buildConstraintClauses(constraints, litSet)

	// add these clauses to the master list
	clauses = append(clauses, constraintClauses...)

	// now, we can convert the clauses into a formula,
	// then solve that formula.
	formula := cnf.Formula(clauses)

	// with the solved formula, we can find the literals that have the assigned UIDs,
	// and then return the constrainters that contain those elements.
	solution, ok := solveSetCoverFormula(formula, st)

	selectedConstrainers := make([]Constrainer, 0, len(solution))
	for _, elem := range solution {
		if elem.selected {
			selectedConstrainers = append(selectedConstrainers, elem.c)
		}
	}
	return selectedConstrainers, ok
}

// TODO refactor to 1) solve the formula in one func
//                  2) find the results in another func
func solveSetCoverFormula(formula cnf.Formula, st stringTable) ([]*solutionElem, bool) {
	solver := sat.New()
	solver.AddFormula(formula)

	if satisfied := solver.Solve(); satisfied {
		// then, we found a solution!
		solution := make([]*solutionElem, 0)
		satSolution := solver.Assignments()
		for uid, isSet := range satSolution {
			// look up the uid in string table
			element := st.lookup(cnf.NewLitInt(uid))
			// check to make sure this isn't a unit constraint literal
			if element.meta == nil {
				continue
			}
			sol := &solutionElem{
				c:        element.meta.owner.(Constrainer),
				selected: isSet,
			}
			solution = append(solution, sol)
		}

		return solution, true
	}

	return nil, false
}

// this function builds each of the constraint clauses
// from each constraint
func buildConstraintClauses(constraints []Constraint, literals []*setElem) []cnf.Clause {
	clauses := []cnf.Clause{}

	// TODO each of these checks can be done in parallel... turn this into a parfor loop.
	// for each of the constraints, find the subset of metadata elements that satisfy it
	for _, constraint := range constraints {

		// build the clause showing which
		// set elements satisfy this particular constraint
		constraintClause := buildSingleConstraintClause(constraint, literals)

		// finally, add this constraint clause to the list of
		// clauses representing our formula
		clauses = append(clauses, constraintClause)
	}
	return clauses
}

func buildSingleConstraintClause(constraint Constraint, literals []*setElem) cnf.Clause {
	// get the subset of literals, and the ID for each of the sets.
	// this is the subset of items that satisfy the constraint
	var (
		satisfied []*setElem = filterUnsatisfied(constraint, literals)
		// build the clause that consists of each of the satisfied
		// elements + the negated unit claus
		clause = make([]cnf.Lit, 0, len(satisfied)+1)
	)

	// add the negation of the unit constraint literal
	clause = append(clause, cnf.NewLit(constraint.uid, false).Neg())

	// add each of the satified literals
	// to the clause
	for _, elem := range satisfied {
		clause = append(clause, elem.lit)
	}
	return cnf.Clause(clause)
}

// this maps each constraint to a unit clause to represent it
// array -> 'a' :: array is the constraint, 'a' is the UID representing the constraint
func buildUnitConstraints(constraints []Constraint, st stringTable) []cnf.Clause {
	var clauses = make([]cnf.Clause, 0, len(constraints))
	for i, constraint := range constraints {
		lit := cnf.NewLit(constraint.uid, false)
		clauses[i] = cnf.Clause{lit}
		st.add(&setElem{
			lit:  lit,
			meta: nil,
			name: "ConstraintLiteral",
		})
	}
	return clauses
}

// takes the subset of the metadata that
// satisfies the constraint, and keeps track of the indices of those elements
func filterUnsatisfied(constraint Constraint, literals []*setElem) []*setElem {

	var subset = make([]*setElem, 0, len(literals))

	for _, data := range literals {
		if data.meta.Satisfies(constraint) {
			subset = append(subset, data)
		}
	}
	return subset
}

// takes an array of constrainters and collects them into a []s*etElems
func toLitSet(dataSet []Constrainer) []*setElem {
	litSet := make([]*setElem, 0, len(dataSet))
	for i, constrainer := range dataSet {
		meta := constrainer.GetMetadata()
		litSet[i] = &setElem{
			name: meta.Name(),
			meta: meta,
			lit:  cnf.NewLit(UID(), false),
		}
	}
	return litSet
}
