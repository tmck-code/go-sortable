package Value

type Sortable interface { Value() interface{} }

type SortableInt struct { Val int }
type SortableFloat struct { Fval int }

func (l SortableInt) Value() interface{} { return l.Val }
func (l SortableFloat) Value() interface{} { return l.Fval }

func Compare(l, r Sortable) int {
	if l.Value().(int) < r.Value().(int) {
		return -1
	} else if l.Value().(int) == r.Value().(int) {
		return 0
	} else {
		return 1
	}
}

