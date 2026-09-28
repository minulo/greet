// Package greet is a tiny library for showing gopkg-charmed at work.
//
// Import it as gopkg-ps7.pfe.staging.canonical.com/minulo/greet.v2: the ".v2" selects branch v2.
package greet

import "fmt"

// Release names the state of branch v2.
const Release = "2.0"

// Hello returns a greeting for name in the language lang: "en", "fr" or "vi".
//
// Version 2 added lang. That breaks every caller of version 1, so version 2
// has its own import path, and a program can use both side by side.
func Hello(name, lang string) string {
	format := "Hello, %s!"
	switch lang {
	case "fr":
		format = "Bonjour, %s!"
	case "vi":
		format = "Xin chào, %s!"
	}
	return fmt.Sprintf(format+" (greet v2, release %s)", name, Release)
}
