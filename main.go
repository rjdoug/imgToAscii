package main

import (
	"fmt"
	"log"

	"gopkg.in/gographics/imagick.v2/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	err := mw.ReadImage("car.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sucessfully loaded image!")
	height := mw.GetImageHeight()
	width := mw.GetImageWidth()
	fmt.Printf("Image size: %d x %d\n", width, height)

	var imagePixels [480][640][3]float64
	//func (mw *MagickWand) GetImagePixelColor(x, y int) (color *PixelWand, err error)
	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {
			pw, err := mw.GetImagePixelColor(i, j)
			if err != nil {
				log.Fatal(err)
			}
			imagePixels[i][j][0] = pw.GetRed()
			imagePixels[i][j][1] = pw.GetGreen()
			imagePixels[i][j][2] = pw.GetBlue()
			// fmt.Printf("Index: %d,%d\n", i,j)
			// fmt.Printf("R: %f\nG: %f\nB: %f\n", mapToRgb(pw.GetRed()), mapToRgb(pw.GetGreen()), mapToRgb(pw.GetBlue()))
			// fmt.Println()
		}
	}
	fmt.Println("DONE!")
	for i := 0; i < int(height); i++ {
		for j := 0; j < int(width); j++ {
			fmt.Printf("Index: %d,%d\n", i, j)
			fmt.Printf("R: %f\nG: %f\nB: %f\n", mapToRgb(imagePixels[i][j][0]), mapToRgb(imagePixels[i][j][1]), mapToRgb(imagePixels[i][j][2]))
			fmt.Println()
		}
	}

	// 	cols, rows, err := mw.GetSize()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(cols)
	// 	fmt.Println(rows)
}

func mapToRgb(color float64) float64 {
	return color * 255
}


