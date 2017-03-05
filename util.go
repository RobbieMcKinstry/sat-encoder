package encoder

import (
	"fmt"
)

func StringProposition(label, val string) Proposition {
	return func(m *Metadata) bool {
		return m.LabelMatchesString(label, val)
	}
}

func IntProposition(label string, val int) Proposition {
	return func(m *Metadata) bool {
		return m.LabelEqualsInt(label, val)
	}
}

func (l LabelExistsError) Error() string {
	return fmt.Sprintf("The label %s already exists.", l.label)
}

func makeUIDGenerator() func() int {
	var n int = 1
	buf := make(chan int)

	go func() {
		for {
			buf <- n
			n++
		}
	}()

	return func() int {
		return <-buf
	}
}

func ResetUIDGenerator() {
	UID = makeUIDGenerator()
}
