package mediator

import "fmt"

type Mediator interface {
	Notify(event string, data interface{})
}

type Payment struct {
	mediator Mediator
}

func (p *Payment) Pay() error {
	p.mediator.Notify("payment.paid", p)
	return nil
}

type Order struct {
	mediator Mediator
}

func (o *Order) Create() error {
	o.mediator.Notify("order.created", o)
	return nil
}

type Notification struct {
	mediator Mediator
}

func (n *Notification) Send(message string) error {
	fmt.Println("Sending notification: ", message)
	return nil
}

type OrderMediator struct {
	payment      *Payment
	order        *Order
	notification *Notification
}

func NewOrderMediator(
	payment *Payment,
	order *Order,
	notification *Notification,
) *OrderMediator {
	m := &OrderMediator{
		payment:      payment,
		order:        order,
		notification: notification,
	}

	payment.mediator = m
	order.mediator = m
	notification.mediator = m
	return m
}

func (m *OrderMediator) Notify(event string, data interface{}) {
	switch event {
	case "payment.paid":
		m.order.Create()
	case "order.created":
		m.notification.Send("Order created")
	}
}

func Main() {
	mediator := NewOrderMediator(
		&Payment{},
		&Order{},
		&Notification{},
	)

	mediator.payment.Pay()
}
