package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Image path is required")
	}

	imgPath := os.Args[1]

	file, err := os.Open(imgPath)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	greyImg := toGreyScale(img)

	outputFile, err := os.Create("output.jpg")
	if err != nil {
		log.Fatalln(err)
	}

	defer outputFile.Close()

	err = jpeg.Encode(outputFile, greyImg, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func toGreyScale(img image.Image) image.Image {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)

	modifiedImg := image.NewRGBA(rect)

	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			red := float64(originalColor.R)
			green := float64(originalColor.G)
			blue := float64(originalColor.B)

			grey := uint8(math.Round((red + green + blue) / 3))

			modifiedColor := color.RGBA{
				R: grey,
				G: grey,
				B: grey,
				A: originalColor.A,
			}

			modifiedImg.Set(x, y, modifiedColor)
		}
	}
	return modifiedImg
}
