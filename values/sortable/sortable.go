package SortableValues

type Sortable interface {
	Value() interface{}
	Compare(Sortable) int
}
