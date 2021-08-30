# zero

[![pkg.go.dev][gopkg-badge]][gopkg]

`zero` zero finds unnecessary assignment which zero value assigns to a variable.

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
