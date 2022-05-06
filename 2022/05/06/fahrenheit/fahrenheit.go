package fahrenheit

type Fahrenheit struct{}

func New(country string) Fahrenheit {
	return Fahrenheit{}
}

func (f Fahrenheit) Value() float64 {
	return 85.5
}
