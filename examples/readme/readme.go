package main

import (
	"fmt"

	opt "github.com/romnnn/configo"
)

type myAppConfig struct {
	EnableExperimental *opt.Flag
}

var defaultConfig myAppConfig = myAppConfig{
	EnableExperimental: opt.SetFlag(false),
}

func main() {
	userConfig := myAppConfig{
		EnableExperimental: opt.SetFlag(true),
	}
	// Merge user config with the default config
	// Will only set values that have NOT been set
	// Otherwise, use OverrideConfig
	opt.MergeConfig(&userConfig, defaultConfig)
	didEnableExperimental := opt.Enabled(userConfig.EnableExperimental)
	fmt.Printf("EnableExperimental=%t\n", didEnableExperimental) // true
}
