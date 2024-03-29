package SortableContainers

import "github.com/tmck-code/go-sortable/values/sortable"

type SortableList struct {
	items []SortableValues.Sortable
	size  int
}

func (l *SortableList) Sort() {
	for i, value := range(l.items) {
        j := i - 1
        for j >= 0 && l.items[j].Compare(value) == 1 {
            l.items[j+1] = l.items[j]
            j = j - 1
        }
        l.items[j+1] = value
    }
}

func (l *SortableList) Add(el SortableValues.Sortable) {
	l.items = append(l.items, el)
	l.size += 1
}
