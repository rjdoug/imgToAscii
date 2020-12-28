package main

import (
	"fmt"
	"log"
	"strings"
	"math"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	// Split asscii string. These are the characters from lightest to darkest
	asciiString := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	characters := strings.Split(asciiString, "")

	mw := imagick.NewMagickWand()
	err := mw.ReadImage("teArohaClock.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sucessfully loaded image!")

	// func (mw *MagickWand) AdaptiveResizeImage(cols, rows uint) error
	err = mw.AdaptiveResizeImage(500, 200)

	if err != nil {
		log.Fatal(err)
	}
	height := mw.GetImageHeight()
	width := mw.GetImageWidth()
	fmt.Printf("Image size: %d x %d\n", width, height)

	// can't set final in golang, so using very large numbers as array bounds
	var imagePixels [10000][10000]int

	//func (mw *MagickWand) GetImagePixelColor(x, y int) (color *PixelWand, err error)
	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {
			pw, err := mw.GetImagePixelColor(j, i)
			if err != nil {
				log.Fatal(err)
			}
			// imagePixels[i][j] = averageBrightness(mapToRgb(pw.GetRed()), mapToRgb(pw.GetGreen()), mapToRgb(pw.GetBlue()))
			// imagePixels[i][j] = lightness(mapToRgb(pw.GetRed()), mapToRgb(pw.GetGreen()), mapToRgb(pw.GetBlue()))
			imagePixels[i][j] = luminosity(mapToRgb(pw.GetRed()), mapToRgb(pw.GetGreen()), mapToRgb(pw.GetBlue()))
			asciiIndex := math.Floor(float64(imagePixels[i][j]) / 3.93)
			// fmt.Print(characters[int(asciiIndex)])
			// fmt.Print(characters[int(asciiIndex)])
			fmt.Print(characters[int(asciiIndex)])
			
		}
		fmt.Println("")
	}
	fmt.Println("DONE!")
}

func mapToRgb(color float64) float64 {
	return color * 255
}

// Gets a brightness value for the current pixel
func averageBrightness(r, g, b float64) int {
	
	return (int(r) + int(g) + int(b)) / 3
}

func lightness(r, g, b float64) int {
	value := math.Max(math.Max(r, g), b)
	return int(value)
}

func luminosity(r, g, b float64) int {
	value := 0.21 * r + 0.72 * g + 0.07 * b
	return int(value)
}


// Lightness: average the maximum and minimum values out of R, G and B - max(R, G, B) + min(R, G, B) / 2
// Luminosity: take a weighted average of the R, G and B values to account for human perception - 0.21 R + 0.72 G + 0.07 B
