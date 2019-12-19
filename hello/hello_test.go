package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string){
		if got != want{
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Marcel", "French")
		want := "Bonjour, Marcel"

		assertCorrectMessage(t, got, want)
	})
}
