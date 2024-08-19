# clydetools

A simple Go package for learning about upcoming Clyde FC fixtures.

Usage:

```
package main

import (
	"fmt"
	"os"

	"github.com/smacktoid/clydetools"
)

func main() {
	fixtures, err := clydetools.GetFixtures()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	length := len(fixtures)
	fmt.Printf("The next %d Clyde fixtures are\n\n", length)
	for _, fixture := range fixtures {
		fmt.Println(fixture)
	}
}
```

A valid (Football API)[] key must be present on the environment as `CLYDETOOLS_API_KEY` for 