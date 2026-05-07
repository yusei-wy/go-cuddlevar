# cuddlevar

A Go linter that detects and auto-fixes unnecessary blank lines between variable assignments and block statements that use the assigned variable.

## Motivation — Why not WSL?

[WSL](https://github.com/bombsimon/wsl) is a comprehensive whitespace linter for Go with 20+ rules. cuddlevar fills a specific gap that WSL does not cover:

| | WSL | cuddlevar |
|---|---|---|
| Direction | **Permissive** — `allow-whole-block` *permits* cuddling but never *enforces* it | **Enforcement** — reports blank lines that should be removed and provides auto-fix |
| `go func` / `defer func` | Checks if the `go` statement invokes a function assigned on the line above | Checks if the assigned variable is *used inside the closure body* |
| Auto-fix | Varies by rule | `analysis.SuggestedFix` for `go vet -fix` and IDE quick-fix |

WSL's `force-err-cuddling` only enforces cuddling for `err` variables. There is no generic "force cuddle when variable is used in block" option as of WSL v5.8.0.

**cuddlevar is designed to complement WSL, not replace it.**

## Before / After

### Assignment before `if`

```go
// Before
x := 1

if x > 0 {
    println(x)
}

// After
x := 1
if x > 0 {
    println(x)
}
```

### Assignment before `for` / `range`

```go
// Before
items := []int{1, 2, 3}

for _, item := range items {
    println(item)
}

// After
items := []int{1, 2, 3}
for _, item := range items {
    println(item)
}
```

### Assignment before `switch`

```go
// Before
mode := "json"

switch mode {
case "json":
    println("json")
}

// After
mode := "json"
switch mode {
case "json":
    println("json")
}
```

### Assignment before `go func()`

```go
// Before
ch := make(chan int)

go func() {
    ch <- 1
}()

// After
ch := make(chan int)
go func() {
    ch <- 1
}()
```

### Assignment before `defer func()`

```go
// Before
cleanup := func() {}

defer func() {
    cleanup()
}()

// After
cleanup := func() {}
defer func() {
    cleanup()
}()
```

### `var` declaration before block

```go
// Before
var x int

if x > 0 {
    println(x)
}

// After
var x int
if x > 0 {
    println(x)
}
```

### Multi-assignment before block

```go
// Before
x, y := 1, 2

if y > 0 {
    println(y)
}

// After
x, y := 1, 2
if y > 0 {
    println(y)
}
```

## Not Detected

cuddlevar intentionally skips these cases:

- **Variable not used in block** — the blank line may be intentional grouping
- **Already cuddled** — no blank line to remove
- **Comment between assignment and block** — the comment is intentional separation
- **Blank identifier (`_`)** — `_ = expr` is not a meaningful assignment

## Install

### golangci-lint plugin (recommended)

Create `.custom-gcl.yml`:

```yaml
version: v2.12.1
plugins:
  - module: github.com/yusei-wy/go-cuddlevar
    version: v1.0.0
```

Build the custom binary:

```sh
golangci-lint custom
```

Add to `.golangci.yml`:

```yaml
linters:
  enable:
    - cuddlevar
  settings:
    custom:
      cuddlevar:
        type: module
```

Run:

```sh
custom-gcl run
```

### Standalone

```sh
go install github.com/yusei-wy/go-cuddlevar/cmd/cuddlevar@latest
cuddlevar ./...
```

## License

[MIT](LICENSE)
