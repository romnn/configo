package configo

import "testing"

type myConfig struct {
	BoolFlag *Flag
}

type myComplexConfig struct {
	Sanitize           *Flag
	ShowErrors         *Flag
	EnableExperimental *Flag
}

func TestSet(t *testing.T) {
	if (*Set(false)).value {
		t.Errorf("Setting false did yield true")
	}
	if !(*Set(true)).value {
		t.Errorf("Setting true did yield false")
	}
}

func TestEnabled(t *testing.T) {
	trueVal := Set(true)
	falseVal := Set(false)
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

func TestConfigOverride(t *testing.T) {
	o := myConfig{BoolFlag: Set(true)}
	oo := myConfig{BoolFlag: Set(false)}
	OverrideConfig(&o, oo)
	if Enabled(o.BoolFlag) {
		t.Errorf("BoolFlag was not successfully overridden")
	}
}

func TestConfigMerge(t *testing.T) {
	o := myConfig{BoolFlag: Set(true)}
	oo := myConfig{BoolFlag: Set(false)}
	MergeConfig(&o, oo)
	if !Enabled(o.BoolFlag) {
		t.Errorf("BoolFlag was not successfully merged")
	}
}

func TestComplexConfigMerge(t *testing.T) {
	defaultConfig := myComplexConfig{
		EnableExperimental: Set(false),
		Sanitize:           Set(true),
	}
	userConfig := myComplexConfig{
		EnableExperimental: Set(true),
		ShowErrors:         Set(true),
	}
	MergeConfig(&userConfig, defaultConfig)
	if !Enabled(userConfig.Sanitize) {
		t.Errorf("Sanitize was not merged to true")
	}
	if !Enabled(userConfig.ShowErrors) {
		t.Errorf("ShowErrors was not merged to true")
	}
	if !Enabled(userConfig.EnableExperimental) {
		t.Errorf("EnableExperimental was not merged to true")
	}
}
