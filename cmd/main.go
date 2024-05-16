package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/Ezan777/go-image-modifier/effects"
)

func main() {
	fmt.Println("Hello World!")
	reader, err := os.Open("assets/AloySeyka.jpg")

	if err != nil {
		fmt.Println("ERROR LOADING IMAGE: " + err.Error())
	}

	img, _, _ := image.Decode(reader)
	reader.Close()
	
	outImage := effects.ApplyEffect(img)
	f, _ := os.Create("assets/AloySeyka(1).jpg")
	jpeg.Encode(f, outImage, &jpeg.Options{Quality: 100})
	fmt.Println("Image saved")
}