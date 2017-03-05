package encoder

import (
	"github.com/mitchellh/go-sat/cnf"
)

type setElem struct {
	lit  cnf.Lit
	meta *Metadata
}

// TODO make a global string table mapping literals to their names.
// TODO add a func on Constrainer that is called Name()
// Merge the toMetadata and the toSetElem funcs into one...
// it should take the constrainer and derive from it the Name, the UID, and the Metadata
func SetCover(dataSet []Constrainer, constraints []Constraint) []Constrainer {

	metadataSet := toMetadataSet(dataSet)
	litSet := toLitSet(metadataSet)

	// for each of the constraints, find the subset of metadata elements that satisfy it
	for i, constraint := range constraints {
		subset, indices := filterUnsatisfied(constraint, metadataSet)
		literals := []cnf.Lit{}

		// now, generate the clause for that subset.

		// add the constraint literal
		baseLit := cnf.NewLit(constraint.uid, false)
		literals = append(literals, baseLit)

		// create a literal for each of the elements on the subset
		for j, satisfied := range subset {
			_, _, _, _, _ = i, j, satisfied, indices, litSet
			// add a new literal.......
		}
	}
	return nil
}

// takes the subset of the metadata that
// satisfies the constraint, and keeps track of the indices of those elements
func filterUnsatisfied(constraint Constraint, metadataSet []*Metadata) ([]*Metadata, []int) {

	var subset = make([]*Metadata, 0, len(metadataSet))
	var indices = make([]int, 0, len(metadataSet))

	for index, data := range metadataSet {
		if data.Satisfies(constraint) {
			subset = append(subset, data)
			indices = append(indices, index)
		}
	}
	return subset, indices
}

// takes an array of metadata elements, and generates a unique literal for them
func toLitSet(meta []*Metadata) []*setElem {
	var result = make([]*setElem, 0, len(meta))
	for i, metadata := range meta {
		result[i] = &setElem{
			meta: metadata,
			lit:  cnf.NewLit(UID(), false),
		}
	}
	return result
}

// takes an array of constrainers, collects their metadata.
func toMetadataSet(dataSet []Constrainer) []*Metadata {
	metadataSet := make([]*Metadata, 0, len(dataSet))
	for i, constrainer := range dataSet {
		metadataSet[i] = constrainer.GetMetadata()
	}
	return metadataSet
}
