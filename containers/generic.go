package Container

import "github.com/tmck-code/go-sortable/values"

type Reversible interface {
	Add()
	Reverse()
}

type ReversibleList struct {
	size  int
	items []interface{}
}

func (l *ReversibleList) Add(el Value.Generic) {
	l.items = append(l.items, el)
	l.size += 1
}

func (l *ReversibleList) Reverse() {
	for i := len(l.items)/2-1; i >= 0; i-- {
		opp := len(l.items)-1-i
		l.items[i], l.items[opp] = l.items[opp], l.items[i]
	}
}
