package externalapis

type PaymentGateway struct{}

func NewPaymentGateway() *PaymentGateway {
	return &PaymentGateway{}
}

func (g *PaymentGateway) ProcessPayment(amount float64, currency string) error {
	return nil
}
