package main

import (
	"StructuresV2/cmd/internal"
	"fmt"
)

const PURCHASE = 1700

func main() {

	partner := internal.NewPartner("Peter", 25, 14000, 1000)
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	//cust.WrOffDebt()
	startTransaction(cust)
	startTransaction(partner)
	fmt.Printf("%+v\n", cust)
	fmt.Printf("%+v\n", partner)
	cust.ProcessingDiscount()
	g, err := internal.CalcPrice(cust, PURCHASE)
	fmt.Printf("%+v\n", g, err)
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
