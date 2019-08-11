package values

type Sortable interface {
	Value() interface{}
	Compare(Sortable) int
}
