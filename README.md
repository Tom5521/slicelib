# Slicelib

This repository contains a collection of utility functions and types for working with slices in Go. These utilities provide convenient methods for manipulating and comparing slices, making your Go code more efficient and readable.

## Installation

You can install my module with this command

```bash
go get github.com/Tom5521/slicelib
```

## Usage

### Examples

```go
package main

import (
	"fmt"
	"strings"

	"github.com/Tom5521/slicelib"
)

func main(){
	a := slicelib.Slice[string]{"a1", "a2", "a3", "b1", "b2", "b3"}

	a.Filter(func(s string) bool {
		return strings.HasPrefix(s, "a")
	})

	fmt.Println(a) // Output: [a1 a2 a3]
}
```

```go
package main

import (
	"fmt"

	"github.com/Tom5521/slicelib"
)

func main() {
	a := slicelib.Slice[string]{"a1", "a2", "a3", "b1", "b2", "b3"}

	contains := a.Contains("b2")

	fmt.Println(contains) // Output: true
}
```

### The Slice types has the following methods:

- Slice
- Append
- Clear
- Copy
- Index
- Insert
- Delete
- Pop
- Remove
- Reverse
- IsEmpty
- Len
- Contains
- RemoveDuplicates
- Equal
- EqualFunc
- SortFunc
- Filter

#### Only on OrderedSlice:

- BinarySearch
- Sort

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
