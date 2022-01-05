package strategy_demo

import "testing"

func TestExamplePayByCash(t *testing.T) {
	payment := NewPayment("Ada","",99,&Cash{})
	payment.Pay()
}
func TestExamplePayByBank(t *testing.T) {
	payment := NewPayment("Ada","",99,&Bank{})
	payment.Pay()
}
