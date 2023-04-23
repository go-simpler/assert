# assert

[![ci](https://github.com/go-simpler/assert/actions/workflows/go.yml/badge.svg)](https://github.com/go-simpler/assert/actions/workflows/go.yml)
[![docs](https://pkg.go.dev/badge/go-simpler.org/assert.svg)](https://pkg.go.dev/go-simpler.org/assert)
[![report](https://goreportcard.com/badge/go-simpler.org/assert)](https://goreportcard.com/report/go-simpler.org/assert)
[![codecov](https://codecov.io/gh/go-simpler/assert/branch/main/graph/badge.svg)](https://codecov.io/gh/go-simpler/assert)

Common assertions to use with the standard testing package

## 📌 About

`assert` is a minimalistic replacement for the [stretchr/testify][1] package,
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
There is also a special tool to do this automatically,
add the following directive to any `.go` file in the root of your project and run `go generate ./...`:

```go
//go:generate go run -tags=installer go-simpler.org/assert/cmd/installer .
```

See the `cmd/installer` documentation for details.

## 📋 Usage

The `dotimport` subpackage should be dot-imported so that `E` and `F` can be used as local types:

```go
"go-simpler.org/assert"
. "go-simpler.org/assert/dotimport"
```

Optional format and arguments can be provided to any assertion to customize the error message:

```go
// prints "actual 1; expected 2" instead of "got 1; want 2"
assert.Equal[E](t, 1, 2, "actual %d; expected %d", 1, 2)
```

## 🧪 Assertions

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

## ❤️ Credits

Inspired by [matryer/is][2].
The idea of the internal implementation belongs to [xakep666][3].

[1]: https://github.com/stretchr/testify
[2]: https://github.com/matryer/is
[3]: https://github.com/xakep666
