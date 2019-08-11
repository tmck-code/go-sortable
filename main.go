package main

import (
	"./values"
	"./containers"
	"fmt"
)

func main() {
	i := values.Int{Val: 4}
	fmt.Printf("%+v\n", i)

	s := new(containers.SortableList)

	s.Add(values.Float{Fval: 2})
	s.Add(values.Int{Val: 1})
	fmt.Printf("%+v\n", s)

	s.Sort()
	fmt.Printf("%+v\n", s)
}
