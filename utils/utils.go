package utils

import "image"

func NewCanvas(img image.Image) (*image.RGBA, image.Point) {
	size := img.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)

	newCanvas := image.NewRGBA(rect)

	return newCanvas, size
}
