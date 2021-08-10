# noctor

[![pkg.go.dev][gopkg-badge]][gopkg]

`noctor` finds unnecessary constructor like functions.

```go
package a

type T struct {
	N int
	S string
}

func NewT1(n int, s string) *T { // want "unnecessary constructor like function"
	return &T{
		N: n,
		S: s,
	}
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/noctor
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/noctor?status.svg

