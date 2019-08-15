package Container

import "github.com/tmck-code/go-sortable/values"

type Sortable interface {
	Sort()
	Add()
}

type SortableList struct {
	Items []Value.Sortable
	size  int
}

func (l *SortableList) Sort() {
	for i, value := range(l.Items) {
        j := i - 1
        for j >= 0 && Value.Compare(l.Items[j], value) == 1 {
            l.Items[j+1] = l.Items[j]
            j = j - 1
        }
        l.Items[j+1] = value
    }
}

func (l *SortableList) Add(el Value.Sortable) {
	l.Items = append(l.Items, el)
	l.size += 1
}
