package term

import (
	"fmt"
	"os"
)

func ExampleNew() {
	t := New()
	for _, opt := range []Option{Red, Green, Blue} {
		t.Set(opt)
		fmt.Fprintln(t, "Hello, world!")
	}
	// Output: [31mHello, world!
	// [39m[32mHello, world!
	// [39m[34mHello, world!
	// [39m
}

func ExampleNewWriter() {
	w := NewWriter(os.Stdout)
	w.Set(Red)
	fmt.Fprint(w, "Hello, world!")
	// Output: [31mHello, world![39m
}

func ExampleWriter_With() {
	t := New()
	t.Set(Yellow)
	fmt.Fprintf(
		t,
		"The url %s is used for examples",
		t.With(Blue, Underline).Render("https://example.com/"),
	)
	// Output: [33mThe url [4;34mhttps://example.com/[24;33m is used for examples[39m
}
