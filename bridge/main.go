package bridge

import "fmt"

// PaymentGateway interface
// This is the bridge between the PaymentMethod and the PaymentGateway
type PaymentGateway interface {
	Charge(amount float64) error
}

type CreditCard struct{}

func (c *CreditCard) Charge(amount float64) error {
	fmt.Println("Paying by CreditCard $", amount)
	return nil
}

type PayPal struct{}

func (p *PayPal) Charge(amount float64) error {
	fmt.Println("Paying by PayPal $", amount)
	return nil
}

// PaymentMethod interface
// This is the interface that the client will use to pay
type PaymentMethod interface {
	Pay(amount float64) error
}

type BasePayment struct {
	gateway PaymentGateway
}

type CreditCardPayment struct {
	BasePayment
}

func (c *CreditCardPayment) Pay(amount float64) error {
	fmt.Println("Pay by CARD")
	return c.gateway.Charge(amount)
}

type PayPalPayment struct {
	BasePayment
}

func (p *PayPalPayment) Pay(amount float64) error {
	fmt.Println("Pay by PayPal")
	return p.gateway.Charge(amount)
}
