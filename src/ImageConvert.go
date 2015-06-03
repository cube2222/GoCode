package main
import "image"
import (
	"image/jpeg"
	"os"
	"fmt"
	"image/png"
)

func main() {
	var ThisImage image.Image

	in, err := os.Open("D:\\output.jpg")
	defer in.Close()
	if err != nil {
		fmt.Println(err)
	}

	ThisImage, err = jpeg.Decode(in)
	if err != nil {
		fmt.Println(err)
	}

	out, err := os.Create("D:\\output.png")
	if err != nil {
		fmt.Println(err)
	}

	png.Encode(out, ThisImage)
	out.Close()
}
