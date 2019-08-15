package main

import (
	"fmt"
	"github.com/tmck-code/go-sortable/containers"
	"github.com/tmck-code/go-sortable/values"
)

func main() {
	// Create a new "sortable" container
	// Objects only need to implement the Value() and Compare() methods,
	// in order to provide a sortable value to the container
	// This method works with things like sets - could define Hash()
	s := new(Container.SortableList)

	// Add our "sortable" values to the container
	s.Add(Value.SortableFloat{Fval: 2})
	s.Add(Value.SortableInt{Val: 1})
	// Check that everything was added to the container
	fmt.Printf("original: %+v\n", s)

	// Sort the contents!
	s.Sort()
	fmt.Printf("sorted:   %+v\n", s)
}
