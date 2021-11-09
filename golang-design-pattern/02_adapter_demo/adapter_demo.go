package adapter_demo
type Target interface {
	Request() string
}
type Adaptee interface {
	SpecificRequest() string
}
type adapteeImpl struct {

}

func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func NewAdaptee() Adaptee{
	return &adapteeImpl{}
}
//Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}
//Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
//NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target{
	return &adapter{
		Adaptee: adaptee,
	}
}