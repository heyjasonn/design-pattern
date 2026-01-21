package facade

import "fmt"

type PaymentService struct {
}

func (p *PaymentService) Pay(amount float64) error {
	fmt.Println("Paying $", amount)
	return nil
}

type Gmail struct{}

func (g *Gmail) SendEmail(to string, subject string, body string) error {
	fmt.Println("Sending email to ", to, " with subject ", subject, " and body ", body)
	return nil
}

type OrderService struct {
}

func (o *OrderService) CreateOrder(amount float64) error {
	fmt.Println("Creating order for $", amount)
	return nil
}

type OrderFacade struct {
	orderService   *OrderService
	paymentService *PaymentService
	gmail          *Gmail
}

func (f *OrderFacade) CreateOrder(amount float64) error {
	f.orderService.CreateOrder(amount)
	f.paymentService.Pay(amount)
	f.gmail.SendEmail("test@test.com", "Order Created", "Order created successfully")
	return nil
}

func Main() {
	orderFacade := &OrderFacade{
		orderService:   &OrderService{},
		paymentService: &PaymentService{},
		gmail:          &Gmail{},
	}

	orderFacade.CreateOrder(100)
}
