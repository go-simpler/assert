# assert

[![checks](https://github.com/go-simpler/assert/actions/workflows/checks.yml/badge.svg)](https://github.com/go-simpler/assert/actions/workflows/checks.yml)
[![pkg.go.dev](https://pkg.go.dev/badge/go-simpler.org/assert.svg)](https://pkg.go.dev/go-simpler.org/assert)
[![goreportcard](https://goreportcard.com/badge/go-simpler.org/assert)](https://goreportcard.com/report/go-simpler.org/assert)
[![codecov](https://codecov.io/gh/go-simpler/assert/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/assert)

Common assertions to use with the `testing` package.

## 📌 About

`assert` is a minimalistic replacement for the `stretchr/testify` package,
providing an alternative API to switch between `t.Errorf()` and `t.Fatalf()`:
instead of using separate packages (`assert`/`require`), it ~~ab~~uses type parameters:

```go
assert.Equal[E](t, 1, 2) // [E] for t.Errorf()
assert.Equal[F](t, 1, 2) // [F] for t.Fatalf()
```

## 📦 Install

```shell
go get go-simpler.org/assert
```

> [!note]
> This package is not even meant to be a dependency!
> It's tiny (<100 LoC), so you can just copy-paste it into your project.
> There is also a special tool to do this automatically:
> just add the following directive to any `.go` file and run `go generate ./...`:
> ```go
> //go:generate go run -tags=copier go-simpler.org/assert/cmd/copier@latest
> ```
> See the `cmd/copier` documentation for details.

## 📋 Usage

The `EF` subpackage should be dot-imported so that `E` and `F` can be used as local types:

```go
"go-simpler.org/assert"
. "go-simpler.org/assert/EF"
```

Optional format and arguments can be provided to any assertion to customize the error message:

```go
assert.Equal[E](t, 1, 2, "%d != %d", 1, 2) // prints "1 != 2"
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
