# term

Simple helpers for formatting terminal output

## Usage

```
w := term.New()
w.Set(term.Yellow)
link := w.With(term.Blue, term.Underline)
fmt.Fprintf(w, "The url %s is used for examples\n", link.Render("https://example.com/"))
```

## Installation

```
$ go get github.com/madss/term
```

## License

MIT

## Author

Mads Sejersen (a.k.a madss on github)
