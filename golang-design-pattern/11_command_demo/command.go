package command_demo

import "fmt"

type Command interface {
	Execute()
}
type StartCommand struct {
	mb *MotherBoard
}
type MotherBoard struct{}
func NewStartCommand(mb *MotherBoard) *StartCommand{
	return &StartCommand{mb: mb}
}
func (c *StartCommand) Execute(){
	c.mb.Start()
}
type RebootCommand struct {
	mb *MotherBoard
}
func NewRebootCommand(mb *MotherBoard) *RebootCommand{
	return &RebootCommand{mb:mb}
}
func (c *RebootCommand) Execute(){
	c.mb.Reboot()
}
func (mb *MotherBoard) Start() {
	fmt.Print("system starting\n")
}
func (mb *MotherBoard) Reboot() {
	fmt.Print("system rebooting\n")
}
type Box struct {
	button1 Command
	button2 Command
}
func NewBox(button1,button2 Command) *Box{
	return  &Box{button1: button1,button2: button2}
}
func (b *Box) PressButton1(){
	b.button1.Execute()
}
func (b *Box) PressButton2(){
	b.button2.Execute()
}