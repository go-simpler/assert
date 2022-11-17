# assert

[![ci](https://github.com/junk1tm/assert/actions/workflows/go.yml/badge.svg)](https://github.com/junk1tm/assert/actions/workflows/go.yml)
[![docs](https://pkg.go.dev/badge/github.com/junk1tm/assert.svg)](https://pkg.go.dev/github.com/junk1tm/assert)
[![report](https://goreportcard.com/badge/github.com/junk1tm/assert)](https://goreportcard.com/report/github.com/junk1tm/assert)
[![codecov](https://codecov.io/gh/junk1tm/assert/branch/main/graph/badge.svg)](https://codecov.io/gh/junk1tm/assert)

Common assertions to use with the standard testing package

## 📌 About

`assert` is a minimalistic replacement for the famous [testify][1] package,
providing an alternative syntax for switching between `t.Errorf()` and
`t.Fatalf()`: instead of using separate packages (`assert`/`require`), it
~~ab~~uses type parameters:

```go
assert.Equal[E](t, 1, 2) // [E] for t.Errorf()
assert.Equal[F](t, 1, 2) // [F] for t.Fatalf()
```

## 📦 Install

```shell
go get github.com/junk1tm/assert
```

⚠️ This package is not even meant to be a dependency! It's tiny (<100 LoC) and
[MIT-licensed](LICENSE), so you can just copy-paste it into your project. There
is a special tool to do this automatically, just add the following directive to
any `.go` file in the root of your project and run `go generate ./...`:

```go
//go:generate go run -tags=installer github.com/junk1tm/assert/cmd/installer .
```

See `cmd/installer` documentation for details.

## 📋 Usage

One should import the `dotimport` subpackage using dot form to be able to use
`E`/`F` parameters as local types:

```go
"github.com/junk1tm/assert"
. "github.com/junk1tm/assert/dotimport"
```

Optional format and arguments can be provided to any assertion to customize the
error message:

```go
// prints "actual 1; expected 2" instead of "got %1; want %2"
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

Inspired by [matryer/is][2]. The idea of the internal implementation belongs
to [xakep666][3].

[1]: https://github.com/stretchr/testify
[2]: https://github.com/matryer/is
[3]: https://github.com/xakep666
