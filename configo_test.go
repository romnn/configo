package configo

import "testing"

type myConfig struct {
	BoolFlag *Flag
}

type myComplexConfig struct {
	Sanitize           *Flag
	ShowErrors         *Flag
	EnableExperimental *Flag
	BatchSize          *Int // Example where 0 value makes no sense
}

func TestConfigOverride(t *testing.T) {
	o := myConfig{BoolFlag: SetFlag(true)}
	oo := myConfig{BoolFlag: SetFlag(false)}
	OverrideConfig(&o, oo)
	if Enabled(o.BoolFlag) {
		t.Errorf("BoolFlag was not successfully overridden")
	}
}

func TestConfigMerge(t *testing.T) {
	o := myConfig{BoolFlag: SetFlag(true)}
	oo := myConfig{BoolFlag: SetFlag(false)}
	MergeConfig(&o, oo)
	if !Enabled(o.BoolFlag) {
		t.Errorf("BoolFlag was not successfully merged")
	}
}

func TestComplexConfigMerge(t *testing.T) {
	defaultConfig := myComplexConfig{
		EnableExperimental: SetFlag(false),
		Sanitize:           SetFlag(true),
		BatchSize:          SetInt(64),
	}
	userConfig := myComplexConfig{
		EnableExperimental: SetFlag(true),
		ShowErrors:         SetFlag(true),
		BatchSize:          SetInt(32),
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
	if GetIntOrDefault(userConfig.BatchSize, 0) != 32 {
		t.Errorf("EnableExperimental was not merged to true")
	}
}
