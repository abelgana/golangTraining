package facotry

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)

	if err != nil {
		t.Error("Payment method creation failed")
	} else {
		msg := payment.Pay(10.0)
		if !strings.Contains(msg, "Payed using Cash") {
			t.Error("Expected payment using cash but used ", msg)
		}
		t.Log(msg)
	}
}

func TestCreatePaymentMethodDebit(t *testing.T) {
	payment, err := GetPaymentMethod(DebitCard)

	if err != nil {
		t.Error("Payment method creation failed")
	} else {
		msg := payment.Pay(10.0)
		if !strings.Contains(msg, "Payed using Debit Card") {
			t.Error("Expected payment using debit card but used ", msg)
		}
		t.Log(msg)
	}
}

func TestCreatePaymentMethodNotSupported(t *testing.T) {
	_, err := GetPaymentMethod(20)

	if err == nil {
		t.Error("This payment Method shouldn't exist")
	}
	t.Log(err)
}
