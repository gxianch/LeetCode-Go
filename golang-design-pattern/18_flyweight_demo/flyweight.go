package _8_flyweight_demo

import "fmt"

type ImageFlyweight struct {
	data string
}
type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}
var imageFactory *ImageFlyweightFactory
func GetImageFlyweightFactory() *ImageFlyweightFactory{
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps :make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}
func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight{
	image := f.maps[filename]
	if image == nil{
		image = NewImageFlyweight(filename)
		f.maps[filename]=image
	}
	return image
}
func NewImageFlyweight(filename string) *ImageFlyweight{
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}
func (i *ImageFlyweight) Data() string{
	return i.data
}
type ImageViewer struct {
	*ImageFlyweight
}
func NewImageViewer(filename string) *ImageViewer{
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight:image,
	}
}
func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}