package adapter

import "fmt"

type Payment interface {
	Pay(amount float64) error
}

type CreditCardAdapter struct {
	creditCard *CreditCard
}

func (c *CreditCardAdapter) Pay(amount float64) error {
	return c.creditCard.Pay(amount)
}

type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) error {
	fmt.Println("Paying by CreditCard $", amount)
	return nil
}
