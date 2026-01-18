package factory

func Main() {
	factory := &PayPal{}
	payment := factory.CreatePayment()
	payment.Pay(10)

	factory2 := &CreditCard{}
	payment2 := factory2.CreatePayment()
	payment2.Pay(20)
}
