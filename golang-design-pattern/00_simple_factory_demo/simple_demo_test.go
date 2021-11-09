package simple_factory_demo

import "testing"

func TestTypeDemo1(t *testing.T){
	api :=NewAPI(1)
	s := api.Say("Tom")
	if s!="Hi,Tom"{
		t.Fatal("Type1 test fail")
	}
}
func TestTypeDemo2(t *testing.T){
	api :=NewAPI(2)
	s := api.Say("Tom")
	if s!="Hello,Tom"{
		t.Fatal("Type2 test fail")
	}
}