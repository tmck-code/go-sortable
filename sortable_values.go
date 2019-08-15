package main

import "fmt"

type Sortable interface {
	Value() interface{}
}

// Int
type Int struct {
	Val int
}

// Float
type Float struct {
	Fval int
}

func (l Int) Value() interface{} {
	return l.Val
}

func (l Float) Value() interface{} {
	return l.Fval
}

func Compare(l, r Sortable) int {
	if l.Value().(int) < r.Value().(int) {
		return -1
	} else if l.Value().(int) == r.Value().(int) {
		return 0
	} else {
		return 1
	}
}

type SortableContainer interface {
	Sort()
	Add()
}

type SortableList struct {
	items []Sortable
	size  int
}

func (l *SortableList) Sort() {
	for i, value := range(l.items) {
        j := i - 1
        for j >= 0 && Compare(l.items[j], value) == 1 {
            l.items[j+1] = l.items[j]
            j = j - 1
        }
        l.items[j+1] = value
    }
}

func (l *SortableList) Add(el Sortable) {
	l.items = append(l.items, el)
	l.size += 1
}

func main() {
	// Create a new "sortable" container
	// Objects only need to implement the Value() method,
	// in order to provide a sortable value to the container
	// This method works with things like sets - could define Hash()
	s := new(SortableList)

	// Add our "sortable" values to the container
	s.Add(Float{Fval: 2})
	s.Add(Int{Val: 1})
	// Check that everything was added to the container
	fmt.Printf("original: %+v\n", s)

	// Sort the contents!
	s.Sort()
	fmt.Printf("sorted:   %+v\n", s)
}
