package containers

import "../values"

type SortableList struct {
	items []values.Sortable
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

func (l *SortableList) Add(el values.Sortable) {
	l.items = append(l.items, el)
	l.size += 1
}
