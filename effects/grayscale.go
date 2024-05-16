package effects

import (
	"image"
	"image/color"
)

func ApplyEffect(givenImage image.Image) (modifiedImage image.Image) {
	width := givenImage.Bounds().Max.X
	height := givenImage.Bounds().Max.Y

	tempImage := image.NewRGBA(givenImage.Bounds())

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			tempImage.Set(i, j, getAverageBrightness(givenImage.At(i, j)))
		}
	}
	
	return tempImage
}

func getAverageBrightness(initialColor color.Color) (filteredColor color.Color) {
	r, g, b, a := initialColor.RGBA()
	r = r * 0xFF / a
    g = g * 0xFF / a
    b = b * 0xFF / a
	average := (r + g + b)/3
	return color.RGBA{uint8(average), uint8(average), uint8(average), uint8(a)}
}