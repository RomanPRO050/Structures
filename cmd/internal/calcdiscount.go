package internal

import "errors"

const DEFAULT_DISCOUNT = 500

func (c *Customer) ProcessingDiscount() (result int, error error) {

	if !c.discount {
		return 0, errors.New("Discount not available")
	}
	result = DEFAULT_DISCOUNT - c.debt
	if result < 0 {
		return 0, nil
	}
	return result, nil

}
