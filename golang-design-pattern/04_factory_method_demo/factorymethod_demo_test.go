package factory_method_demo

import "testing"

func compute(factory OperatorFactory,a,b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}
func TestOperatorDemo(t *testing.T){
	var (factory OperatorFactory)
	factory=PlusOperatorFactory{}
	if compute(factory ,1 ,2 )!=3{
		t.Fatal("error plus")
	}
	factory = MinusOperatorFactory{}
	if compute(factory,4,2) !=2{
		t.Fatal("error minus")
	}
}