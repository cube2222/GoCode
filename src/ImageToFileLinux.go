package main

import "image/color"
import (
	"image"
	"image/jpeg"
	"os"
)

type MyImage struct {

}
func (i MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (MyImage) At(x, y int) color.Color {
	if y != 0 {
		return color.RGBA{0, 0, uint8(x*x*y*y), 255}
	} else {
		return color.RGBA{0, 0, uint8(x*x*y*y), 255}
	}
}

func (i MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 1024, 1024)
}

func main() {
	var myImage MyImage
	file, err := os.Create("/home/user/myImg.jpg")
	if err != nil {
		panic(err)
	}
	opt := jpeg.Options{}
	opt.Quality = 100
	jpeg.Encode(file, myImage, &opt)
}
