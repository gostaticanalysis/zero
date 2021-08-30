# zero

[![pkg.go.dev][gopkg-badge]][gopkg]

`zero` finds unnecessary assignment which zero value assigns to a variable.

```go
package a

var _ string = "" // want "shoud not assign zero value"

func _() {
	n := 0 // want "shoud not assign zero value"
	_ = n

	var _ []int = nil     // want "shoud not assign zero value"
	var _ []int = []int{} // OK
	m := int32(0)         // OK
	_ = m
	var _ *int = nil            // want "shoud not assign zero value"
	var _ struct{} = struct{}{} // want "shoud not assign zero value"
	var _, _ int                // OK
	var _, _ int = 0, 1         // want "shoud not assign zero value"
	var _, _ int = 1, 2         // OK
	var _, _ int = 1 - 1, 2 - 2 // want "shoud not assign zero value" "shoud not assign zero value"
	var _ bool = false          // want "shoud not assign zero value"
	var _ bool = true           // OK

	type T struct{ N int }
	var _ T = T{} // want "shoud not assign zero value"
}
```

## Install

You can get `zero` by `go install` command (Go 1.16 and higher).

```bash
$ go install github.com/gostaticanalysis/zero/cmd/zero@latest
```

## How to use

`zero` run with `go vet` as below when Go is 1.12 and higher.

```bash
$ go vet -vettool=$(which zero) ./...
```

## Analyze with golang.org/x/tools/go/analysis

You can use [zero.Analyzer](https://pkg.go.dev/github.com/gostaticanalysis/zero/#Analyzer) with [unitchecker](https://golang.org/x/tools/go/analysis/unitchecker).

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gostaticanalysis/zero
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gostaticanalysis/zero?status.svg
