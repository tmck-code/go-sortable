package GenericContainers

import "github.com/tmck-code/go-sortable/values/generic"

type ReversibleList struct {
	size  int
	items []interface{}
}

func (l *ReversibleList) Add(el GenericValues.GenericValue) {
	l.items = append(l.items, el)
	l.size += 1
}

func (l *ReversibleList) Reverse() {
	l.items = l.items
}
