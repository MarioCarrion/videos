package producer1

type Names interface {
	Produce() string
}

func New() Names {
	return str{}
}

type str struct{}

func (f str) Produce() string {
	return "mario"
}

func (c str) String() string {
	return ""
}
