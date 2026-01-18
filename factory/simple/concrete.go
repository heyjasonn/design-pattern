package simple

import "github.com/heyjasonn/design-pattern/factory"

func NewPayment(paymentType string) factory.Payment {
	switch paymentType {
	case "paypal":
		return &factory.PayPal{}
	case "creditcard":
		return &factory.CreditCard{}
	default:
		return nil
	}
}
