package values

type Int struct {
	Val int
}

func (l Int) Value() interface{} {
	return l.Val
}

func (l Int) Compare(r Sortable) int {
	if l.Value().(int) < r.Value().(int) {
		return -1
	} else if l.Value().(int) == r.Value().(int) {
		return 0
	} else {
		return 1
	}
}

