package modes

import (
	"image"
	"image/color"

	"github.com/conur-floki/imageBit/utils"
)

func Negative(img image.Image) image.Image {
	modifiedImg, size := utils.NewCanvas(img)

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			pixelColor := color.RGBA{
				R: 255 - originalColor.R,
				G: 255 - originalColor.G,
				B: 255 - originalColor.B,
				A: 255 - originalColor.A,
			}
			modifiedImg.Set(x, y, pixelColor)
		}
	}
	return modifiedImg
}
