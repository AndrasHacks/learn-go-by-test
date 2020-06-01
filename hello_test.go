package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("Given empty string, when Hello(), then prints Hello, World!", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World!"

		assertEquals(got, want, t)
	})

	t.Run("Given a name, when Hello(), then greets name.", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris!"

		assertEquals(got, want, t)
	})

}

func assertEquals(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
