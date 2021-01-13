# gfycatid

This is a rewrite of [this library](https://github.com/ilsiepotamus/gfycatid)
that generates IDs in the style of [gfycats](https://gfycat.com). I have not
checked up on how socially acceptable / memeworthy gfycats is at the current
moment, so this might be tone-deaf, I'm not sure.

In case you're new to these IDs; they are cute, but you must do
uniqueness-checking on your own.

## "Important" details

- You can generate IDs with any number of adjectives.
- Generators are safe for use with multiple goroutines.
- This library uses static lists that are published in the source files; the
  lists that gfycat publishes online haven't been changed in years.
- You can override those lists by reassigning them at runtime: They're in the
  `assets` subpackage.

## Example

```go
package main

import (
	"fmt"
	"github.com/MichaelDiBernardo/gfycatid"
)

func main() {
	// Create a new generator.
	nadj := 3
	gen := gfycatid.Create(nadj)
	id := gen.Gen()
	fmt.Printf("ID with %d adjectives is: '%s'", nadj, id)
}
```
