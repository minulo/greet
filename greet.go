// Package greet is a tiny library for showing gopkg-charmed at work.
//
// Import it as gopkg-ps7.pfe.staging.canonical.com/minulo/greet.v1: the ".v1" selects branch v1.
package greet

import "fmt"

// Release names the state of branch v1.
const Release = "1.1"

// Hello returns a greeting for name.
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s! (greet v1, release %s)", name, Release)
}
