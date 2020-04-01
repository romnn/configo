package configo

// Flag ...
type Flag struct {
	value bool
}

// SetFlag ...
func SetFlag(enable bool) *Flag {
	return &Flag{value: enable}
}

// Enabled ...
func Enabled(option *Flag) bool {
	if option != nil {
		return (*option).value
	}
	return false
}

// FlagSet ...
func FlagSet(option *Flag) bool {
	return option != nil
}
