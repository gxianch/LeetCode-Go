package _2_bridge_demo

import "fmt"

type AbstractMessage interface {
	SendMessage(text,to string)
}
type MessageImplementer interface {
	Send(text, to string)
}
type MessageSMS struct{}
func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}
func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}
type MessageEmail struct{}
func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}
func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}
type CommonMessage struct {
	method MessageImplementer
}
func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}
func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}
