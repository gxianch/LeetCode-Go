package factory_method_demo
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}
type OperatorFactory interface {
	Create() Operator
}
type OperatorBase struct {
	a,b int
}
func (o *OperatorBase) SetA(a int){
	o.a=a
}
func (o *OperatorBase) SetB(b int){
	o.b=b
}
type PlusOperatorFactory struct {

}
type PlusOperator struct {
	*OperatorBase
}

func (o PlusOperator) Result() int {
	return o.a + o.b
}

func (PlusOperatorFactory) Create() Operator{
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}
type MinusOperatorFactory struct {
	
}
type MinOperator struct {
	 *OperatorBase
}


func (o MinOperator) Result() int {
	return o.a - o.b
}

func (MinusOperatorFactory) Create() Operator{
	return &MinOperator{
		OperatorBase: &OperatorBase{},
	}
}