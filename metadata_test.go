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
