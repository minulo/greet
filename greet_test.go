package greet

import "testing"

func TestHello(t *testing.T) {
	for lang, want := range map[string]string{
		"en": "Hello, Juju! (greet v2, release 2.0)",
		"fr": "Bonjour, Juju! (greet v2, release 2.0)",
		"vi": "Xin chào, Juju! (greet v2, release 2.0)",
	} {
		if got := Hello("Juju", lang); got != want {
			t.Errorf("Hello(%q) = %q, want %q", lang, got, want)
		}
	}
}
