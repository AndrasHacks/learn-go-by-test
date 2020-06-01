package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("")
	want := "Hello Gophers!"

	if got != want {
		printError(got, want, t)
	}
}

func Test_GivenName_WhenHello_GreetsName(t *testing.T) {
	got := Hello("Chris")
	want := "Hello Chris!"

	if got != want {
		printError(got, want, t)
	}
}

func printError(got string, want string, t *testing.T) {
	t.Errorf("Got %q want %q", got, want)
}
