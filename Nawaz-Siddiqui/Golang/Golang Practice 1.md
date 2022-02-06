# Golang Practice 1

**Hello world**

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello world this is nawaz")
}
```

println comes under fmt package.

**Prints a random number**

```go
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favourite number is:", rand.Intn(100))
}
```

[rand.Int](http://rand.Int)n is used to generate random integer with an upper limit of hundred. The function comes under math/rand package.

By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement `package rand`

**Exported names**

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The value of pi is: ", math.Pi)
}
```

n Go, a name is exported if it begins with a capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the `math` package.

`pizza` and `pi` do not start with a capital letter, so they are not exported.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

The above example generates the value of pi in math. Using pi instead of Pi will lead to error.

**Functions**

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("The number is:", add(1, 12))
}
```

A function can take zero or more arguments.

In this example, `add` takes two parameters of type `int`.

Notice that the type comes *after* the variable name.

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

For example:

```
x int, y int
```

to

```
x, y int
```

**Returning multiple results**

A function can return any number of results.

The `swap` function returns two strings.

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```