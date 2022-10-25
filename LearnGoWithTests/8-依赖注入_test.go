package main

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_YiLai(t *testing.T) {
	t.Run("testGreet", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet1(&buffer, "Chris")

		got := buffer.String()
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
func Greet1(writer *bytes.Buffer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
