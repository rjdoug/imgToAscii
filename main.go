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
	err := mw.ReadImage("smiley.png")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sucessfully loaded image!")
	height := mw.GetImageHeight()
	width := mw.GetImageWidth()
	fmt.Printf("Image size: %d x %d\n", width, height)

	var imagePixels [10000][10000]int
	//func (mw *MagickWand) GetImagePixelColor(x, y int) (color *PixelWand, err error)
	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {
			pw, err := mw.GetImagePixelColor(i, j)
			if err != nil {
				log.Fatal(err)
			}
			imagePixels[i][j] = averageBrightness(mapToRgb(pw.GetRed()), mapToRgb(pw.GetGreen()), mapToRgb(pw.GetBlue()))
			asciiIndex := math.Floor(float64(imagePixels[i][j]) / 3.92)
			fmt.Print(characters[int(asciiIndex)])
			fmt.Print(characters[int(asciiIndex)])
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


// Lightness: average the maximum and minimum values out of R, G and B - max(R, G, B) + min(R, G, B) / 2
// Luminosity: take a weighted average of the R, G and B values to account for human perception - 0.21 R + 0.72 G + 0.07 B
