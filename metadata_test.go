package encoder_test

import (
	"github.com/RobbieMcKinstry/sat-encoder"
	"testing"
)

func TestIsInt(t *testing.T) {
	const num int = 5
	const str string = "hello"

	// unhappy path.
	if _, ok := encoder.IsInt(str); ok {
		t.Error("Passed in a string, but IsInt returned true.")
	}

	// happy path
	if i, ok := encoder.IsInt(num); ok {
		if i != num {
			t.Errorf("Passed in %v, got back %v", num, i)
		}
	} else {
		t.Error("Passed in an int, but IsInt returned false.")
	}
}

func TestIsString(t *testing.T) {
	const num int = 5
	const str string = "hello"

	// unhappy path.
	if _, ok := encoder.IsString(num); ok {
		t.Error("Passed in an int, but IsString returned true.")
	}

	// happy path
	if i, ok := encoder.IsString(str); ok {
		if i != str {
			t.Errorf("Passed in %v, got back %v", str, i)
		}
	} else {
		t.Error("Passed in an string, but IsString returned false.")
	}
}

func TestMatchesString(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value string = "bar"
	const bad string = "incorrect***value"

	m.AddLabel(label, value)
	if !m.LabelMatchesString(label, value) {
		t.Error("Added label %v with value %v, but when checking, label did not match.", label, value)
	}

	if m.LabelMatchesString(label, bad) {
		t.Error("Added label %v with value %v, but when checking, label matched %v", label, value, bad)
	}
}

func TestMatchesInt(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value int = 100
	const bad int = 99

	m.AddLabel(label, value)
	if !m.LabelEqualsInt(label, value) {
		t.Error("Added label %v with value %v, but when checking, label did not match.", label, value)
	}

	if m.LabelEqualsInt(label, bad) {
		t.Error("Added label %v with value %v, but when checking, label matched %v", label, value, bad)
	}
}

func TestLtInt(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value int = 100
	const good int = 101
	const bad int = 99

	m.AddLabel(label, value)
	if !m.LabelLessThanInt(label, good) {
		t.Error("Added label %v with value %v, but when checking, %v:%v < %v was false", label, value, label, value, good)
	}

	if m.LabelLessThanInt(label, bad) {
		t.Error("Added label %v with value %v, but when checking, %v:%v < %v was true", label, value, label, value, bad)
	}
}

func TestGtInt(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value int = 100
	const good int = 99
	const bad int = 101

	m.AddLabel(label, value)
	if !m.LabelGreaterThanInt(label, good) {
		t.Error("Added label %v with value %v, but when checking, %v:%v > %v was false", label, value, label, value, good)
	}

	if m.LabelGreaterThanInt(label, bad) {
		t.Error("Added label %v with value %v, but when checking, %v:%v > %v was true", label, value, label, value, bad)
	}
}

func TestGtEqInt(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value int = 100
	const good int = 99
	const bad int = 101

	m.AddLabel(label, value)
	if !m.LabelGEqInt(label, good) {
		t.Error("Added label %v with value %v, but when checking, %v:%v > %v was false", label, value, label, value, good)
	}

	if m.LabelGEqInt(label, bad) {
		t.Error("Added label %v with value %v, but when checking, %v:%v > %v was true", label, value, label, value, bad)
	}
}

func TestLtEqInt(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value int = 100
	const good int = 101
	const bad int = 99

	m.AddLabel(label, value)
	if !m.LabelLEqInt(label, good) {
		t.Error("Added label %v with value %v, but when checking, %v:%v < %v was false", label, value, label, value, good)
	}

	if m.LabelLEqInt(label, bad) {
		t.Error("Added label %v with value %v, but when checking, %v:%v < %v was true", label, value, label, value, bad)
	}
}

func TestNEqInt(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value int = 100
	const bad int = 99

	m.AddLabel(label, value)
	if m.LabelNEqInt(label, value) {
		t.Error("Added label %v with value %v, but when checking, label didn't match %v.", label, value, value)
	}

	if !m.LabelNEqInt(label, bad) {
		t.Error("Added label %v with value %v, but when checking, label matched %v", label, value, bad)
	}
}

func TestLabelExistsError(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value int = 100

	m.AddLabel(label, value)
	err := m.AddLabel(label, value)
	if err == nil {
		t.Error("Expected an error, received no error.")
	}
}

func TestLabelMatches(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value int = 100
	const bad = "bar"

	m.AddLabel(label, value)
	if !m.LabelMatches(label, value) {
		t.Error("Added label %v with value %v, but when checking, label didn't match %v", label, value, value)
	}

	if m.LabelMatches(label, bad) {
		t.Error("Added label %v with value %v, but when checking, label surprisingly matched %v", label, value, bad)
	}
}

func TestLabelMatchesWith(t *testing.T) {
	m := encoder.NewMetadata(nil)
	const label string = "foo"
	const value int = 100
	const bad = "bar"

	matcher := func(a, b interface{}) bool {
		return a == b
	}

	m.AddLabel(label, value)
	if !m.LabelMatchesWith(label, value, matcher) {
		t.Error("Added label %v with value %v, but when checking, label didn't match %v", label, value, value)
	}

	if m.LabelMatches(label, matcher) {
		t.Error("Added label %v with value %v, but when checking, label surprisingly matched %v", label, value, bad)
	}
}
