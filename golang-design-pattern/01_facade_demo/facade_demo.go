package facade_demo

import "fmt"

type  API interface {
	Test() string
}
//AModuleAPI ...
type AModuleAPI interface {
	TestA() string
}
//BModuleAPI ...
type BModuleAPI interface {
	TestB() string
}
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}
type aModuleImpl struct{}
func (*aModuleImpl) TestA() string {
	return "A module running"
}
func NewAModuleAPI() AModuleAPI{
	return &aModuleImpl{}
}
type bModuleImpl struct{}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
//NewBModuleAPI return new BModuleAPI
func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}
func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}
func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

