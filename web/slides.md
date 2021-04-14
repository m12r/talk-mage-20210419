class: center, middle

<img src="/assets/gary.svg" style="height: 33.3vh; width: auto;">

# Mage Build Tool

Talk at GoGraz April 2021 Meetup

---

class: middle

# About Me

Matthias Endler - Independent Software Developer

- Programming since the mid-1980s, professionally since 1998
- Interested in too many things
    - Gardening and Homesteading
    - Model Railroading (FREMO, MiniMax)
    - Programming Languages (too many to list them all)
    - VoIP, RTC, FreeSWITCH and other weird telecom stuff
    - Web Technologies
- Go since 1.8 (2017)
- Currently available for freelance work

---

class: middle

# About Mage

- a make-like build tool created by [Nate Finch](https://npf.io)
- uses Go functions as targets
- has no dependencies except for the Go runtime
- can be bootstraped via Go modules without installation
- can be used to build a *»specialized«* build tool
- <https://magefile.org>
- <https://github.com/magefile/mage>

---

class: middle

# Targets

```go
//+build mage

package main

// Build builds something
func Build() {}

// Install installs something
func Install(ctx context.Context) error { return nil }

// Run runs what
func Run(what string) error { return nil }

// Exec executes something
func Exec(
    ctx context.Context,
    name string,
    count int,
    debug bool,
    timeout time.Duration,
) { return nil }
```

```shell
mage build
mage install
mage run blah
mage exec blubb 42 true 10s

```

---

class: middle

# Aliases and Default

```go
//+build mage

package main

var Default = All

var Aliases = map[string]interface{}{
    "i":     Install,
    "build": Compile,
}

func Build() {}

func Install() {}

func All() {}

```

---

class: middle

# Dependencies

```go
import (
    "fmt"

    "github.com/magefile/mage/mg"
)

func A() {
    mg.Deps(D)

    fmt.Println("A")
}

func B() {
    mg.Deps(mg.F(C, 42), D)

    fmt.Println("B")
}

func C(meaningOfLife int) {
    fmt.Println("C")
}

func D() {
    fmt.Println("D")
}

```

---

class: middle

# Dependencies cont'd

```go
// A and B run in seperate go routines
mg.Deps(A, B)

// same as Deps, but with a specific context
mg.CtxDeps(ctx, A, B)

// A and B are run in serial
mg.SerialDeps(A, B)

// same as SerialDeps, but with a specific context
mg.SerialCtxDeps(ctx, A, B)

```

---

class: middle

# Namespaces

```go
//+build mage

package main

import (
    "github.com/magefile/mage/mg"
)

type Web mg.Namespace

func (Web) Install() {}

func (Web) Build() {}
```

---

class: middle

# Importing Targets

```go
//+build mage

package main

import (
    // mage:import
    _ "example.com/x/foo"

    // mage:import bar
    "example.com/x/foo"

    "github.com/magefile/mage/mg"
)

func TargetA() {
    mg.Deps(foo.MeaningOfLife)
}
```

```go
package foo

import "fmt"

func MeaningOfLife() {
    fmt.Println("42")
}

```

---

class: middle

# Zero Install Option

Just create a go file in the `main` package with the following
content:

```go
//+build ignore

package main

import (
    "os"

    "github.com/magefile/mage/mage"
)

func main() {
    os.Exit(mage.Main())
}
```

We assume, we saved it as `bootstrap.go`, then we can run it
like this:

```shell
go run bootstrap.go
```

---

class: middle

# Compile a specialized build tool

If you need something, which can run without go, then you can
compile your magefile into an executable.

```shell
mage -compile ./specialized-build-tool
```

**Notice:** `mage` ignores `GOOS` and `GOARCH` environment variables,
so if you want to cross-compile, you need to pass it via parameters
`-goos` and `-goarch`.

```shell
mage -compile -goos windows -goarch amd64 ./specialized-build-tool
```

---

class: middle

# Some command line options

## -l

lists the available targets

## -h [target]

shows the long comment of the specified target

## -t [duration]

provides a context, which cancels mage, if it runs longer than
the provided duration.

---

class: middle

# Live coding... 😅

---

class: middle

# Thank you!
