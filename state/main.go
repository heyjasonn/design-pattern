package state

import "fmt"

type Order struct {
	state State
}

type State interface {
	Pay(o *Order) error
	Ship(o *Order) error
	Cancel(o *Order) error
}

func NewOrder() *Order {
	o := &Order{}
	o.SetState(&CreateState{})
	return o
}

func (o *Order) SetState(state State) {
	o.state = state
}

func (o *Order) Pay() error {
	return o.state.Pay(o)
}

func (o *Order) Ship() error {
	return o.state.Ship(o)
}

func (o *Order) Cancel() error {
	return o.state.Cancel(o)
}

type CreateState struct{}

func (c *CreateState) Pay(o *Order) error {
	fmt.Println("Order paid")
	o.SetState(&PaidState{})
	return nil
}

func (c *CreateState) Ship(o *Order) error {
	return fmt.Errorf("Cannot ship order that is not paid")
}

func (c *CreateState) Cancel(o *Order) error {
	fmt.Println("cancelled order")
	o.SetState(&CancelledState{})
	return nil
}

type PaidState struct{}

func (p *PaidState) Pay(o *Order) error {
	return fmt.Errorf("order already paid")
}

func (p *PaidState) Ship(o *Order) error {
	fmt.Println("Shipping order")
	o.SetState(&ShippedState{})
	return nil
}

func (p *PaidState) Cancel(o *Order) error {
	fmt.Println("Canceling order")
	o.SetState(&CancelledState{})
	return nil
}

type ShippedState struct{}

func (s *ShippedState) Pay(o *Order) error {
	return fmt.Errorf("order already shipped")
}

func (s *ShippedState) Ship(o *Order) error {
	return fmt.Errorf("order already shipped")
}

func (s *ShippedState) Cancel(o *Order) error {
	fmt.Println("Cannot cancel shipped order")
	return fmt.Errorf("Cannot cancel shipped order")
}

type CancelledState struct{}

func (c *CancelledState) Pay(o *Order) error {
	return fmt.Errorf("order cancelled")
}

func (c *CancelledState) Ship(o *Order) error {
	return fmt.Errorf("order cancelled")
}

func (c *CancelledState) Cancel(o *Order) error {
	fmt.Println("Cannot cancel cancelled order")
	return fmt.Errorf("Cannot cancel cancelled order")
}

func Main() {
	order := NewOrder()

	order.Pay()
	order.Cancel()
	order.Ship()
}
