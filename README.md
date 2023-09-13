# assert

[![checks](https://github.com/go-simpler/assert/actions/workflows/checks.yml/badge.svg)](https://github.com/go-simpler/assert/actions/workflows/checks.yml)
[![pkg.go.dev](https://pkg.go.dev/badge/go-simpler.org/assert.svg)](https://pkg.go.dev/go-simpler.org/assert)
[![goreportcard](https://goreportcard.com/badge/go-simpler.org/assert)](https://goreportcard.com/report/go-simpler.org/assert)
[![codecov](https://codecov.io/gh/go-simpler/assert/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/assert)

Common assertions to use with the standard testing package

## 📌 About

`assert` is a minimalistic replacement for the [`stretchr/testify`][1] package,
providing an alternative syntax to switch between `t.Errorf()` and `t.Fatalf()`:
instead of using separate packages (`assert`/`require`), it ~~ab~~uses type parameters:

```go
assert.Equal[E](t, 1, 2) // [E] for t.Errorf()
assert.Equal[F](t, 1, 2) // [F] for t.Fatalf()
```

## 📦 Install

```shell
go get go-simpler.org/assert
```

⚠️ This package is not even meant to be a dependency!
It's tiny (<100 LoC), so you can just copy-paste it into your project.
There is also a special tool to do this automatically:
just add the following directive to any `.go` file and run `go generate ./...`:

```go
//go:generate go run -tags=copier go-simpler.org/assert/cmd/copier .
```

See the `cmd/copier` documentation for details.

## 📋 Usage

The `dotimport` subpackage should be dot-imported so that `E` and `F` can be used as local types:

```go
"go-simpler.org/assert"
. "go-simpler.org/assert/dotimport"
```

Optional format and arguments can be provided to any assertion to customize the error message:

```go
// prints "actual/expected" instead of "got/want".
assert.Equal[E](t, 1, 2, "values are not equal\nactual: %v\nexpected: %v", 1, 2)
```

### Equal

Asserts that two values are equal.

```go
assert.Equal[E](t, 1, 2)
```

### NoErr

Asserts that the error is nil.

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

### Panics

Asserts that the given function panics with the argument `v`.

```go
assert.Panics[E](t, func() { /* panic? */ }, 42)
```

## ❤️ Credits

Inspired by [`matryer/is`][2].
The idea of the internal implementation belongs to [xakep666][3].

[1]: https://github.com/stretchr/testify
[2]: https://github.com/matryer/is
[3]: https://github.com/xakep666
