package main

import (
	"flag"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if len(os.Args) < 2 {
		log.Fatalln("Image path is required")
	}

	imgPath := flag.String("file", "", "The path to the file")

	flag.Parse()

	if parsed := flag.Parsed(); !parsed {
		infoLog.Printf("Error opening file: %s", *imgPath)
	}

	file, err := os.Open(*imgPath)
	if err != nil {
		errorLog.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		errorLog.Fatal(err)
	}

	infoLog.Printf("Converting img to greyscale")
	greyImg := toGreyScale(img)

	outputFile, err := os.Create("output.jpg")
	if err != nil {
		errorLog.Fatal(err)
	}

	defer outputFile.Close()

	err = jpeg.Encode(outputFile, greyImg, nil)
	if err != nil {
		errorLog.Fatal(err)
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
