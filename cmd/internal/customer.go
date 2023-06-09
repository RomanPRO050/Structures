package internal

type Customer struct {
	Name         string
	Age          int
	Balance      int
	Debt         int
	Discount     bool
	CalcDiscount func() (int, error)
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Name:     name,
		Age:      age,
		Balance:  balance,
		Debt:     debt,
		Discount: discount,
	}
}

func CalcPrice(c *Customer, PURCHASE int) (int, error) {
	a, err := c.CalcDiscount()
	var b int
	if err == nil {
		b = c.Balance - a - PURCHASE
	}
	if err != nil {
		b = c.Balance - PURCHASE
	}
	return b, err
}
