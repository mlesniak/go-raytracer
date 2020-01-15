package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	nx := 200
	ny := 100

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2

			ir := uint8(255.99 * r)
			ig := uint8(255.99 * g)
			ib := uint8(255.99 * b)

			img.Set(i, ny-j, color.RGBA{ir, ig, ib, 255})
		}
	}
	file, err := os.Create("demo.png")
	must(err)
	must(png.Encode(file, img))
}

// go run main.go|pnmtopng >demo.png && open demo.png
func writePPM(nx int, ny int) {
	// Header
	fmt.Println("P3")
	fmt.Println(nx, ny)
	fmt.Println(255)

	// Pixel
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.2

			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)

			fmt.Println(ir, ig, ib)
		}
	}
}
