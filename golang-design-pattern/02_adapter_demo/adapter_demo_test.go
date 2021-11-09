package adapter_demo

import "testing"
var expect = "adaptee method"

func TestAdapterDemo(t *testing.T)  {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()
	println(res)
	if res != expect{
		t.Fatalf("expect: %s, actual: %s", expect, res)
	}
}
