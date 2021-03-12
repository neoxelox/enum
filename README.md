# enum ![Integration](https://github.com/Neoxelox/enum/workflows/Integration/badge.svg)

**📑 `Golang Enumerators` 📑**

## What

Enum is a package that provides simple enumerators for Go, with IDE autocompletion and any type support. It may not be the best syntax, but it's so useful.

## Install

**`go get github.com/neoxelox/enum`**

## Usage

**`Create Enum`**

```go
package main

import (
	"fmt"

	"github.com/neoxelox/zeus/pkg/enum"
)

type State string

type enumStates = struct {
	COMMITTED   State  // You can use any type you want, for example int.
	IN_PROGRESS State
	DONE        State
	enum.Enum
}

var States = enum.New(&enumStates{
	COMMITTED:   "COMMITTED",
	IN_PROGRESS: "BLOCKED",
	DONE:        "DONE",
}).(*enumStates)

func main() {
	fmt.Println(States.DONE)
}
```

Creating a custom type, from a primitive type, for the enum fields (as in the example above), will provide _lite type assertion_, that is, `States.COMMITTED == "COMMITTED"` will evaluate to `true`. If you want full type assertion, you can create `type State struct{ string }`, and use that type as the type for the enum fields `COMMITTED: State{"COMMITTED"}`. Now you can't compare `States.COMMITTED == "COMMITTED"`. However, you will need to create your own `String`, `Marshalling`, `Text`... methods, to deal with serialization correctly.

**`Check if Alias is in Enum`**

```go
States.Is("COMMITTED")  // true
States.Is("COMPLETED")  // false
States.Is("BLOCKED")  // false
```

**`Get Enum Aliases`**

```go
States.Aliases()  // [COMMITTED IN_PROGRESS DONE]
```

**`Check if Value is in Enum`**

```go
States.Is("COMMITTED")  // true
States.Is("BLOCKED")  // true
States.Is("IN_PROGRESS")  // false
States.Is("COMPLETED")  // false
```

**`Get Enum Values`**

```go
States.Values()  // [COMMITTED BLOCKED DONE]
```

See [`GoDev`](https://pkg.go.dev/github.com/neoxelox/enum) for further documentation.

## Contribute

Feel free to contribute to this project : ) .

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT) - read the [LICENSE](LICENSE) file for details.
