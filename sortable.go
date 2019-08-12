package main

import (
	"fmt"
	"github.com/tmck-code/go-sortable/containers/sortable"
	"github.com/tmck-code/go-sortable/values/sortable"
)

func main() {
	s := new(SortableContainers.SortableList)

	s.Add(SortableValues.Float{Fval: 2})
	s.Add(SortableValues.Int{Val: 1})
	fmt.Printf("original: %+v\n", s)

	s.Sort()
	fmt.Printf("sorted:   %+v\n", s)
}
