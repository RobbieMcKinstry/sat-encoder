package encoder

import (
	"github.com/mitchellh/go-sat/cnf"
)

// adds the metadata's name to the string table.
func (st stringTable) add(elem *setElem) {
	st[elem.lit] = elem
}

func (st stringTable) lookup(lit cnf.Lit) *setElem {
	return st[lit]
}

func (st stringTable) addLiterals(set []*setElem) {
	for _, elem := range set {
		st.add(elem)
	}
}
