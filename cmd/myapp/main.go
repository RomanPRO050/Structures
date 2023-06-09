package main

import (
	"Structures/cmd/internal"
	"errors"
	"fmt"
)

const DEFAULT_DISCOUNT = 500
const PURCHASE = 170

func main() {
	cust := internal.NewCustomer("Dmitry", 23, 10000, 100, false)

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("Discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}
	g, err := internal.CalcPrice(cust, PURCHASE)
	fmt.Printf("%+v\n", g, err)
}
