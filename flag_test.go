package configo

import "testing"

func TestSetFlag(t *testing.T) {
	if (*SetFlag(false)).value {
		t.Errorf("Setting false did yield true")
	}
	if !(*SetFlag(true)).value {
		t.Errorf("Setting true did yield false")
	}
}

func TestEnabled(t *testing.T) {
	trueVal := SetFlag(true)
	falseVal := SetFlag(false)
	if Enabled(falseVal) {
		t.Errorf("Checking ptr to false did yield true")
	}
	if !Enabled(trueVal) {
		t.Errorf("Checking ptr to true did yield false")
	}
	if Enabled(nil) {
		t.Errorf("Checking nil ptr did yield true")
	}
}
