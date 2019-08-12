package main

import (
	"fmt"
	"github.com/tmck-code/go-sortable/containers/generic"
	"github.com/tmck-code/go-sortable/values/generic"
)

type CustomVal interface {
	MyMethod() bool
}

type CustomMap struct {
	mapping map[string]string
	name string
	hash string
}

type Custom struct {
	mapping CustomMap
	id int
}

func main() {
	sl := new(GenericContainers.ReversibleList)
	sl.Add(GenericValues.Int{Val: 1})
	sl.Add(GenericValues.Int{Val: 100})

	c := Custom{
		mapping: CustomMap{
			name: "my thing",
			hash: "obj hash",
			mapping: map[string]string{"key": "val"},
		},
		id: 123,
	}
	sl.Add(c)
	fmt.Printf("original: %+v\n", sl)
	sl.Reverse()
	fmt.Printf("reversed: %+v\n", sl)
}
