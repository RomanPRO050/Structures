package main

import (
	"StructuresV2/cmd/internal"
	"fmt"
)

const PURCHASE = 1700

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 1000, true)

	cust.WrOffDebt()

	fmt.Printf("%+v\n", cust)
	cust.ProcessingDiscount()
	/*	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}*/
	g, err := internal.CalcPrice(cust, PURCHASE)
	fmt.Printf("%+v\n", g, err)
}
