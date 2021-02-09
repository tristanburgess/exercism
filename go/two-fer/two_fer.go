/*
Package twofer provides a library with a single method
ShareWith which takes in a string and outputs a silly message.
*/
package twofer

import "fmt"

// ShareWith takes in a string and outputs a string
// of the form "One for {string}, one for me.".
// If the provided string is empty, the output string is
// of the form "One for you, one for me."
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
