package Money

import "github.com/mohammadrahimi/Inventory_Service_Go/src/Domain/Stock/Errors"


type Money struct {
	amount   float64
	currency string
}

func New(amount float64, currency string) (Money, error) {
	if amount <= 0 && currency == "" {
		return Money{}, Errors.ErrorMoney
	}

	return Money{
		amount:   amount,
		currency: currency,
	}, nil

}
