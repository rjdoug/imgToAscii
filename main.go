package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"golang.org/x/term"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	// Check if an image path is provided as a command-line argument
	if len(os.Args) < 2 {
		fmt.Println("Please provide an image path as an argument.")
		fmt.Println("Usage: go run main.go <image_path>")
		return
	}

	imagePath := os.Args[1]

	// Split ascii string. These are the characters from lightest to darkest
	asciiString := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
	characters := strings.Split(asciiString, "")

	mw := imagick.NewMagickWand()
	err := mw.ReadImage(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully loaded image!")

	// Get terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default to 80 if unable to get terminal width
	}

	// Calculate adaptive height
	originalWidth := mw.GetImageWidth()
	originalHeight := mw.GetImageHeight()
	aspectRatio := float64(originalHeight) / float64(originalWidth)
	newWidth := uint(width)
	newHeight := uint(float64(newWidth) * aspectRatio * 0.5) // Multiply by 0.5 to account for terminal line height

	err = mw.AdaptiveResizeImage(newWidth, newHeight)
	if err != nil {
		log.Fatal(err)
	}

	height := mw.GetImageHeight()
	width = int(mw.GetImageWidth())
	fmt.Printf("ASCII Art size: %d x %d\n", width, height)

	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {
			pw, err := mw.GetImagePixelColor(j, i)
			if err != nil {
				log.Fatal(err)
			}
			luminosityValue := luminosity(mapToRgb(pw.GetRed()), mapToRgb(pw.GetGreen()), mapToRgb(pw.GetBlue()))
			asciiIndex := math.Floor(float64(luminosityValue) / 3.93)
			fmt.Print(characters[int(asciiIndex)])
		}
		fmt.Println("")
	}
	fmt.Println("DONE!")
}

func mapToRgb(color float64) float64 {
	return color * 255
}

func luminosity(r, g, b float64) int {
	value := 0.21*r + 0.72*g + 0.07*b
	return int(value)
}
