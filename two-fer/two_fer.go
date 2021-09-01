// Package twofer contains the ShareWith function.
// https://golang.org/doc/effective_go.html#commentary
package twofer

import "fmt"

// ShareWith returns "One for X, one for me.", where X is the value of the parameter "name".
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
