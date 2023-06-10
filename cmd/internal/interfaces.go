package internal

type Debtor interface {
	WrOffDebt() error
}

type Discounter interface {
	ProcessingDiscount() (int, error, *Customer)
}
