package constants

const LAST_TRANSACTIONS_NUMBER_BY_DEFAULT = 10
const LIMIT_TRANSACTION_REPORT = 1_000_000

const ( // Types of customers
	Individual = iota
	Organization
)

const ( // Type of accounts
	SavingAccout = iota
	CurrentAccount
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
