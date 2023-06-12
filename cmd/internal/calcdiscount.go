package internal

import "errors"

const DEFAULT_DISCOUNT = 500

func (c *Customer) ProcessingDiscount() (result int, error error, k *Customer) {

	if !c.Discount {
		return 0, errors.New("Discount not available"), c
	}
	result = DEFAULT_DISCOUNT - c.Overduer.Debt
	if result < 0 {
		return 0, nil, c
	}
	return result, nil, c

}
