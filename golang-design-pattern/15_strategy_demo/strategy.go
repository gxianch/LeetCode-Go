package strategy_demo

import "fmt"

type PaymentContext struct {
	Name, CardID string
	Money        int
}
type PaymentStrategy interface {
	Pay(*PaymentContext)
}
type Payment struct {
	context *PaymentContext
	strategy PaymentStrategy
}
func NewPayment(name,cardid string,money int,strategy PaymentStrategy) *Payment{
	return &Payment{context: &PaymentContext{Name: name,CardID: cardid,Money: money},strategy: strategy}
}
func (p *Payment) Pay(){
	p.strategy.Pay(p.context)
}
type Cash struct {

}
func (c *Cash) Pay(ctx *PaymentContext){
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}
type Bank struct {

}
func (b *Bank) Pay(ctx *PaymentContext){
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)
}