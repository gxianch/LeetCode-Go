package mediator_demo

import (
	"fmt"
	"strings"
)

type  CDDriver struct {
	Data string
}
func (c *CDDriver) ReadData(){
	c.Data = "music,image"
	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	GetMediatorInstance().changed(c)
}
var mediator *Mediator
func GetMediatorInstance() *Mediator{
	if mediator == nil{
		mediator = &Mediator{}
	}
	return mediator
}
func (c *CPU) Process(data string){
	sp := strings.Split(data,",")
	c.Sound = sp[0]
	c.Video = sp[1]
	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
	GetMediatorInstance().changed(c)
}
func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard: display %s\n", v.Data)
	GetMediatorInstance().changed(v)
}
func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
	GetMediatorInstance().changed(s)
}
func (m *Mediator) changed(i interface{}){
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	}
}
type Mediator struct {
	CD    *CDDriver
	CPU   *CPU
	Video *VideoCard
	Sound *SoundCard
}
type CPU struct {
	Video string
	Sound string
}
type VideoCard struct {
	Data string
}
type SoundCard struct {
	Data string
}