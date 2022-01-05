package _8_flyweight_demo

import "testing"
func TestExampleFlyweight(t *testing.T) {
//func ExampleFlyweight() {
	viewer := NewImageViewer("image1.png")
	viewer.Display()
	// Output:
	// Display: image data image1.png
}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")
	viewer1.Display()
	viewer2.Display()
	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		t.Fail()
	}
}