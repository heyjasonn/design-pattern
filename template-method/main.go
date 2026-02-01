package templatemethod

import "fmt"

type PaymentTemplate interface {
	Validate() error
	Pay(amount float64) error
	AfterPay() error
}

type PaymentProcessor struct {
	template PaymentTemplate
}

func (p *PaymentProcessor) Process(amount float64) error {
	if err := p.template.Validate(); err != nil {
		return err
	}
	if err := p.template.Pay(amount); err != nil {
		return err
	}
	return p.template.AfterPay()
}

type CreditCardPayment struct{}

func (c *CreditCardPayment) Validate() error {
	fmt.Println("Validating credit card payment")
	return nil
}

func (c *CreditCardPayment) Pay(amount float64) error {
	fmt.Println("Paying with credit card $", amount)
	return nil
}

func (c *CreditCardPayment) AfterPay() error {
	fmt.Println("Sending confirmation email")
	return nil
}

func Main() {
	creditCardPayment := &CreditCardPayment{}
	paymentProcessor := &PaymentProcessor{template: creditCardPayment}
	paymentProcessor.Process(100)
}
