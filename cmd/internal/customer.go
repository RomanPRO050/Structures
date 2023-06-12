package internal

import "errors"

type Customer struct {
	Name         string
	Age          int
	Discount     bool
	Overduer     Overduer
	CalcDiscount func() (int, error)
}

func (c *Customer) WrOffDebt() error {
	if c.Overduer.Debt >= c.Overduer.Balance {
		return errors.New("Не возможно списать")
	}
	c.Overduer.Balance -= c.Overduer.Debt
	c.Overduer.Debt = 0
	return nil
}

/*
	func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
		return &Customer{
			Name:     name,
			Age:      age,
			balance:  balance,
			debt:     debt,
			discount: discount,
		}
	}
*/
func CalcPrice(c Discounter, PURCHASE int) (int, error) {
	d, err, k := c.ProcessingDiscount()
	var b int
	if err == nil {
		b = k.Overduer.Balance - (PURCHASE - d)
	}
	if err != nil {
		b = d - PURCHASE
	}
	return b, err
}
