# assert

[![ci](https://github.com/go-simpler/assert/actions/workflows/go.yml/badge.svg)](https://github.com/go-simpler/assert/actions/workflows/go.yml)
[![docs](https://pkg.go.dev/badge/github.com/go-simpler/assert.svg)](https://pkg.go.dev/github.com/go-simpler/assert)
[![report](https://goreportcard.com/badge/github.com/go-simpler/assert)](https://goreportcard.com/report/github.com/go-simpler/assert)
[![codecov](https://codecov.io/gh/go-simpler/assert/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/assert)

Common assertions to use with the standard testing package

## ๐ About

`assert` is a minimalistic replacement for the famous [testify][1] package,
providing an alternative syntax for switching between `t.Errorf()` and
`t.Fatalf()`: instead of using separate packages (`assert`/`require`), it
~~ab~~uses type parameters:

```go
assert.Equal[E](t, 1, 2) // [E] for t.Errorf()
assert.Equal[F](t, 1, 2) // [F] for t.Fatalf()
```

## ๐ฆ Install

```shell
go get github.com/go-simpler/assert
```

โ ๏ธ This package is not even meant to be a dependency! It's tiny (<100 LoC) and
[MIT-licensed](LICENSE), so you can just copy-paste it into your project. There
is a special tool to do this automatically, just add the following directive to
any `.go` file in the root of your project and run `go generate ./...`:

```go
//go:generate go run -tags=installer github.com/go-simpler/assert/cmd/installer .
```

See the `cmd/installer` documentation for details.

## ๐ Usage

The `dotimport` subpackage should be dot-imported, so the `E`/`F` parameters
could be used as local types:

```go
"github.com/go-simpler/assert"
. "github.com/go-simpler/assert/dotimport"
```

Optional format and arguments can be provided to any assertion to customize the
error message:

```go
// prints "actual 1; expected 2" instead of "got %1; want %2"
assert.Equal[E](t, 1, 2, "actual %d; expected %d", 1, 2)
```

## ๐งช Assertions

### Equal

Asserts that two values are equal.

```go
assert.Equal[E](t, 1, 2)
```

### NoErr

Asserts that `err` is nil.

```go
assert.NoErr[E](t, err)
```

### IsErr

Asserts that `errors.Is(err, target)` is true.

```go
assert.IsErr[E](t, err, os.ErrNotExist)
```

### AsErr

Asserts that `errors.As(err, target)` is true.

```go
assert.AsErr[E](t, err, new(*os.PathError))
```

## โค๏ธ Credits

Inspired by [matryer/is][2]. The idea of the internal implementation belongs
to [xakep666][3].

[1]: https://github.com/stretchr/testify
[2]: https://github.com/matryer/is
[3]: https://github.com/xakep666
