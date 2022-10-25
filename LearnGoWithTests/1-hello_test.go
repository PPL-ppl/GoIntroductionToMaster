package main

import (
	"testing"
)

func TestHello(t *testing.T) {

	//assert.Equal(t, expected, actual)
	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}

	t.Run("saying hello to Chris", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	})
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

const englishHelloPrefix = "Hello " //常量定义
const frenchHelloPrefix = "Hole "
const spanishHelloPrefix = "Spanish "

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
