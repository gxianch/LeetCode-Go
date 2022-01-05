package _0_decorator_demo

type Component interface {
	Calc() int
}
type ConcreteComponent struct {

}

func (c *ConcreteComponent) Calc() int{
	return 0
}
type MulDecorator struct {
	Component
	num int
}
func WarpMulDecorator(c Component,num int) Component{
	return &MulDecorator{
		Component: c,
		num:num,
	}
}
func (d *MulDecorator) Cal() int{
	return d.Component.Calc()*d.num
}
type AddDecorator struct {
	Component
	num int
}
func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num: num,
	}
}
func (d *AddDecorator) Calc() int{
	return d.Component.Calc()+d.num
}