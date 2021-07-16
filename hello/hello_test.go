package main

import "testing"

// Basic Test Hello
func TestHello(t *testing.T) {
	got := Hello("Bob")
	want := "Hello, Bob"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHelloUsingSubtests(t *testing.T) {
	// testing.TB is an interface that both *testing.T and *testing.B satisfy
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()	// needed to tell the test suite that this method is a helper
		// By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bob")
		want := "Hello, Bob"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}