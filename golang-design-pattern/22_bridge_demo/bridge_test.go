package _2_bridge_demo

import "testing"

func TestExampleCommonSMS(t *testing.T) {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
	// Output:
	// send have a drink? to bob via SMS
}
