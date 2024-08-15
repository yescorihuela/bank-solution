package constants

const LAST_TRANSACTIONS_NUMBER_BY_DEFAULT = 10

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
