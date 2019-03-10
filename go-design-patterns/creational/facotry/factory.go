package facotry

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) string
}

type CashPM struct {
}

type DebitPM struct {
}

const (
	Cash = iota
	DebitCard
)

func GetPaymentMethod(pm int) (PaymentMethod, error) {
	switch pm {
	case Cash:
		return new(CashPM), nil
	case DebitCard:
		return new(DebitPM), nil
	default:
		return nil, errors.New("Payment method not supporeted")
	}
}

func (c *CashPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f Payed using Cash", amount)
}

func (c *DebitPM) Pay(amount float64) string {
	return fmt.Sprintf("%0.2f Payed using Debit Card", amount)
}
