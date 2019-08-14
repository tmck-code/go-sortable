package main

import (
	"fmt"
	"github.com/tmck-code/go-sortable/containers/generic"
	"github.com/tmck-code/go-sortable/values/generic"
)

// Let's define our own custom struct that we want to use
type CustomMap struct {
	mapping map[string]string
	name string
	hash string
}

func main() {
	// Create a new "generic" container. It doesn't care what you store in it.
	sl := new(GenericContainers.ReversibleList)

	// Add some "generic" values from this lib
	sl.Add(GenericValues.Int{Val: 1})
	sl.Add(GenericValues.Int{Val: 100})

	// Add our own custom struct
	sl.Add(CustomMap{
		name: "my thing",
		hash: "obj hash",
		mapping: map[string]string{"key": "val"},
	})
	// Check that everything was added to the container
	fmt.Printf("original: %+v\n", sl)

	// Now, reverse the contents of the list!
	// We don't need to implement anything for this to happen.
	sl.Reverse()
	fmt.Printf("reversed: %+v\n", sl)
}
