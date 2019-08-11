package main

import (
	"github.com/tmck-code/go-sortable/values/sortable"
	"github.com/tmck-code/go-sortable/containers/sortable"
	"fmt"
)

func main() {
	i := SortableValues.Int{Val: 4}
	fmt.Printf("%+v\n", i)

	s := new(SortableContainers.SortableList)

	s.Add(SortableValues.Float{Fval: 2})
	s.Add(SortableValues.Int{Val: 1})
	fmt.Printf("%+v\n", s)

	s.Sort()
	fmt.Printf("%+v\n", s)
}
