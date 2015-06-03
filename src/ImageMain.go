package main
import "image/color"
import "image"
import (
	"image/jpeg"
	"os"
	"fmt"
	"math"
)

type MyImage struct {
}

func (i MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 255, 255)
}

func (i MyImage) At(x, y int) color.Color {
	MyColor := color.RGBA{0, 0, uint8(math.Abs(float64(x - y))), 255}
	return MyColor
}

func main() {
	var ThisImage MyImage
	out, err := os.Create("D:\\output.jpg")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var opt jpeg.Options
	opt.Quality = 100

	err = jpeg.Encode(out, ThisImage, &opt)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
