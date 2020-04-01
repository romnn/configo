package configo

// Int ...
type Int struct {
	value int
}

// SetInt ...
func SetInt(value int) *Int {
	return &Int{value: value}
}

// GetIntOrNil ...
func GetIntOrNil(i *Int) *int {
	if i != nil {
		return &i.value
	}
	return nil
}

// GetIntOrDefault ...
func GetIntOrDefault(i *Int, def int) int {
	if i != nil {
		return i.value
	}
	return def
}
