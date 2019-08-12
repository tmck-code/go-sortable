package GenericValues

type Int struct {
	Val int
}

func (i Int) Value() interface{} {
	return i
}
