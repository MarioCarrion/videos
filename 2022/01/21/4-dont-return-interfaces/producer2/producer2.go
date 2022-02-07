package producer2

type Cars struct{}

func New() Cars {
	return Cars{}
}

func (c Cars) Produce() string {
	return "bmw"
}

func (c Cars) String() string {
	return ""
}
