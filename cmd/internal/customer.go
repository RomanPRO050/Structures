package internal

import "errors"

type Customer struct {
	Name         string
	Age          int
	balance      int
	debt         int
	discount     bool
	CalcDiscount func() (int, error)
}

func (c *Customer) WrOffDebt() error {
	if c.debt >= c.balance {
		return errors.New("Не возможно списать")
	}
	c.balance -= c.debt
	c.debt = 0
	return nil
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		balance:  balance,
		debt:     debt,
		discount: discount,
	}
}

func CalcPrice(c Discounter, PURCHASE int) (int, error) {
	d, err, k := c.ProcessingDiscount()
	var b int
	if err == nil {
		b = k.balance - (PURCHASE - d)
	}
	if err != nil {
		b = d - PURCHASE
	}
	return b, err
}
