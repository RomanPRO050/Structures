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

func CalcPrice(c *Customer, PURCHASE int) (int, error) {
	a, err := c.ProcessingDiscount()
	var b int
	if err == nil {
		b = c.balance - (PURCHASE - a)
	}
	if err != nil {
		b = c.balance - PURCHASE
	}
	return b, err
}
