package flagvar

import "fmt"

// Array represents string array flag variable.
type Array []string

func (a *Array) String() string {
	return fmt.Sprintf("%v", *a)
}

// Set appends element to the array.
func (a *Array) Set(value string) error {
	*a = append(*a, value)
	return nil
}
