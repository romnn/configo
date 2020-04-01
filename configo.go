package configo

import "github.com/imdario/mergo"

// Version is incremented using bump2version
const Version = "0.1.1"

// Config ...
type Config interface{}

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
