package observer

import "fmt"

type Observer interface {
	OnNotify(event string, data any)
}

type Subject interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Notify(event string, data any)
}

type Order struct {
	observers []Observer
}

func (o *Order) Subscribe(observer Observer) {
	o.observers = append(o.observers, observer)
}

func (o *Order) Unsubscribe(observer Observer) {
	for i, observer := range o.observers {
		if observer == observer {
			o.observers = append(o.observers[:i], o.observers[i+1:]...)
		}
	}
}

func (o *Order) Notify(event string, data any) {
	for _, observer := range o.observers {
		observer.OnNotify(event, data)
	}
}

func (o *Order) Create() {
	o.Notify("order.created", "123")
}

type EmailObserver struct {
	email string
}

func (e *EmailObserver) OnNotify(event string, data any) {
	fmt.Println("Sending email to ", e.email, " with event ", event, " and data ", data)
}

type SMSObserver struct {
	phone string
}

func (s *SMSObserver) OnNotify(event string, data any) {
	fmt.Println("Sending SMS to ", s.phone, " with event ", event, " and data ", data)
}

func Main() {
	order := &Order{}
	emailObserver := &EmailObserver{email: "test@test.com"}
	smsObserver := &SMSObserver{phone: "1234567890"}
	order.Subscribe(smsObserver)
	order.Subscribe(emailObserver)
	order.Create()
}
