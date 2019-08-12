package main

import (
	"fmt"
	"github.com/tmck-code/go-sortable/containers/sortable"
	"github.com/tmck-code/go-sortable/values/generic"
	"github.com/tmck-code/go-sortable/values/sortable"
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

	gi := GenericValues.Int{Val: 1}
	fmt.Printf("%+v\n", gi)
}
