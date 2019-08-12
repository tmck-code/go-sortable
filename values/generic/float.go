package GenericValues

type Float struct {
	Val float64
}

func (i Float) Value() interface{} {
	return i
}
