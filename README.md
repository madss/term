# term

[![Go Reference](https://pkg.go.dev/badge/github.com/madss/term.svg)](https://pkg.go.dev/github.com/madss/term) [![Go Report Card](https://goreportcard.com/badge/github.com/madss/term)](https://goreportcard.com/report/github.com/madss/term)

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
