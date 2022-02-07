package sum

type Value struct {
	A int
	B int
}

func Add(a, b Value) Value {
	if a == b {
		return Value{}
	}

	return Value{
		A: a.A + b.A,
		B: b.B + b.B,
	}
}
