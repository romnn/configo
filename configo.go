package configo

import "github.com/imdario/mergo"

// Version is incremented using bump2version
const Version = "0.1.0"

// Config ...
type Config interface{}

// Flag ...
type Flag struct {
	value bool
}

// Set ...
func Set(enable bool) *Flag {
	return &Flag{value: enable}
}

// Enabled ...
func Enabled(option *Flag) bool {
	if option != nil {
		return (*option).value
	}
	return false
}

// OverrideConfig ...
func OverrideConfig(c Config, override Config) {
	if err := mergo.Merge(c, override, mergo.WithOverride); err != nil {
		panic(err)
	}
}

// MergeConfig ...
func MergeConfig(c Config, merge Config) {
	if err := mergo.Merge(c, merge); err != nil {
		panic(err)
	}
}
