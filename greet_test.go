package greet

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, Juju! (greet v1, release " + Release + ")"
	if got := Hello("Juju"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
