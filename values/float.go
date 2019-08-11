package values

type Float struct {
	Fval int
}

func (l Float) Value() interface{} {
	return l.Fval
}

func (l Float) Compare(r Sortable) int {
	if l.Value().(int) < r.Value().(int) {
		return -1
	} else if l.Value().(int) == r.Value().(int) {
		return 0
	} else {
		return 1
	}
}
