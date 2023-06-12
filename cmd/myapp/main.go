package main

import (
	"StructuresV2/cmd/internal"
	"errors"
	"fmt"
	"io"
)

const PURCHASE = 1700

func main() {

	partner := internal.NewPartner("Peter", 25, 14000, 1000)
	//cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)
	custv2Money := internal.Overduer{Balance: 10000, Debt: 1000}
	custv2 := internal.Customer{Name: "Dmitry", Age: 23, Discount: true, Overduer: custv2Money}
	//cust.WrOffDebt()
	startTransaction(&custv2)
	startTransaction(partner)
	startTransactionDynamic(&custv2)
	fmt.Printf("%+v\n", custv2)
	fmt.Printf("%+v\n", partner)
	custv2.ProcessingDiscount()
	g, err := internal.CalcPrice(&custv2, PURCHASE)
	fmt.Printf("%+v\n", g, err)
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}

func startTransactionDynamic(debtor internal.Debtor) error {
	a, ok := debtor.(io.WriteCloser)
	if !ok {
		return errors.New("Incorrect type")
	}
	_, err := a.Write([]byte("Hellom world!"))
	return err
}
