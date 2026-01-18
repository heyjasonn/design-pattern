package factory

import "fmt"

// Paypal payment method
type PayPal struct{}

func (p *PayPal) Pay(amount float64) error {
	fmt.Println("Paying by PayPal $", amount)
	return nil
}

func (p *PayPal) CreatePayment() Payment {
	return &PayPal{}
}

// CreditCard payment method
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) error {
	fmt.Println("Paying by CreditCard $", amount)
	return nil
}

func (c *CreditCard) CreatePayment() Payment {
	return &CreditCard{}
}
