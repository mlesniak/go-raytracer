package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	nx := 200
	ny := 100

	origin := Vector{0.0, 0.0, 0.0}
	lowerLeft := Vector{-2.0, -2.0, -1.0}
	horizontal := Vector{4.0, 0.0, 0.0}
	vertical := Vector{0.0, 2.0, 0.0}

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := Ray{origin, lowerLeft.Add(horizontal.Scale(u).Add(vertical.Scale(v)))}

			c := backgroundColor(r)
			ir := uint8(255.99 * c.R())
			ig := uint8(255.99 * c.G())
			ib := uint8(255.99 * c.B())

			img.Set(i, ny-j, color.RGBA{ir, ig, ib, 255})
		}
	}
	file, err := os.Create("demo.png")
	must(err)
	must(png.Encode(file, img))
}

func backgroundColor(r Ray) Vector {
	unitDirection := Unit(r.Direction())
	t := 0.5 * (unitDirection.Y() + 1.0)
	v1 := Vector{1.0, 1.0, 1.0}.Scale(1.0 - t)
	v2 := Vector{0.5, 0.7, 1.0}.Scale(t)
	return v1.Add(v2)
}
