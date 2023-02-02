package skip

//go:generate stringer -type=Currency

type Currency int

const (
	USD Currency = iota
	MXN
	PHP
)

//go:generate stringer -type=Country

type Country int

const (
	USA Country = iota
	Mexico
	Phillipines
)
