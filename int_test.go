package configo

import "testing"

func TestSetInt(t *testing.T) {
	if (*SetInt(2)).value != 2 {
		t.Errorf("Setting int value did yield wrong result")
	}
	if (*SetInt(0)).value != 0 {
		t.Errorf("Setting int value did yield wrong result")
	}
}

func TestGetOrNil(t *testing.T) {
	if GetIntOrNil(nil) != nil {
		t.Errorf("Checking nil ptr did yield non nil")
	}
	if *GetIntOrNil(SetInt(12)) != 12 {
		t.Errorf("Checking non nil ptr to int did yield wrong value")
	}
	if GetIntOrDefault(nil, 12) != 12 {
		t.Errorf("Checking nil ptr with default did yield wrong value")
	}
	if GetIntOrDefault(SetInt(11), 5) != 11 {
		t.Errorf("Checking non-nil ptr to int with default did yield wrong value")
	}
}
