package main

import (
	"fmt"

	"github.com/MircoT/go-string-fuzzy-finder/pkg/core"
)

// The playground now uses square brackets for type parameters. Otherwise,
// the syntax of type parameter lists matches the one of regular parameter
// lists except that all type parameters must have a name, and the type
// parameter list cannot be empty. The predeclared identifier "any" may be
// used in the position of a type parameter constraint (and only there);
// it indicates that there are no constraints.

func main() {
	finder := core.SimpleFinder{}
	finder.Init()
	finder.SetMinThreshold(0.6)

	result, _ := finder.BestMatch("hLl", []string{"hello", "HELLO", "heaven", "hotel", "heLL", "lol"})

	fmt.Println(result)

	results, _ := finder.Similars("hLlo", []string{"hello", "HELLO", "heaven", "hotel", "heLL", "lol"})

	fmt.Println(results)
}
