package functional

import "github.com/heyjasonn/design-pattern/factory"

func NewPaymentFactory() func() factory.Payment {
	return func() factory.Payment {
		return &factory.PayPal{}
	}
}
