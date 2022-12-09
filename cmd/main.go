package main

import (
	"flag"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/conur-floki/imageBit/modes"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if len(os.Args) < 2 {
		log.Fatalln("Image path is required")
	}

	imgPath := flag.String("file", "", "The path to the file")

	mode := flag.String("mode", "", "Mode to convert image")

	flag.Parse()

	if parsed := flag.Parsed(); !parsed {
		infoLog.Printf("Error parsing the flag: %s", *imgPath)
	}

	file, err := os.Open(*imgPath)
	if err != nil {
		errorLog.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		errorLog.Fatal(err)
	}

	var outputImg image.Image

	switch *mode {
	case "grey":
		infoLog.Printf("Converting img to greyscale")
		outputImg = modes.ToGreyScale(img)
	case "negative":
		infoLog.Printf("Moving Bits")
		outputImg = modes.Negative(img)
	}

	outputFile, err := os.Create("output.jpg")
	if err != nil {
		errorLog.Fatal(err)
	}

	defer outputFile.Close()

	err = jpeg.Encode(outputFile, outputImg, nil)
	if err != nil {
		errorLog.Fatal(err)
	}
}
