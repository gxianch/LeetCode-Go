package facade_demo

import "testing"

var expect = "A module running\nB module running"
func TestFacadeAPIDemo(t *testing.T){
	api :=NewAPI()
	ret :=api.Test()
	if ret != expect {
		t.Fatalf("expect %s, return %s", expect, ret)
	}
}