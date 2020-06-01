package main

import "testing"

func TestHello(t *testing.T) {

	// We might get rid of the *testing.T argument by applying partial application
	assertEquals := func(got string, want string, t *testing.T) {
		t.Helper()
		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	}

	t.Run("Given empty string, when Hello(), then prints Hello, World!", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"

		assertEquals(got, want, t)
	})

	t.Run("Given a name, when Hello(), then greets name.", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris!"

		assertEquals(got, want, t)
	})

	t.Run("Given Spanish language, when Hello(), greets in spanish!", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertEquals(got, want, t)
	})

	t.Run("Given French language, when Hello(), greets in French.", func(t *testing.T) {
		got := Hello("Claude", "French")
		want := "Bonjur, Claude!"
		assertEquals(got, want, t)
	})

}
