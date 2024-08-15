package constants

const (
	Individual = iota
	Organization
)

const (
	USD = iota
	EUR
	CLP
	COP
	MXN
	ARS
	CAD
)

var Currencies = []int{
	USD,
	EUR,
	CLP,
	COP,
	MXN,
	ARS,
	CAD,
}

const (
	WithDrawal = iota
	Deposit
)

const (
	Pending = iota
	Approved
	Rejected
)
