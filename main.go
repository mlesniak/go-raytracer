package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
)

func main() {
	nx := 200
	ny := 100
	ns := 100
	fmt.Printf("Computing %d pixel\n", nx*ny)

	world := World{}
	world.Add(Sphere{Vector{0, 0, -1}, 0.5})
	world.Add(Sphere{Vector{0, -100.5, -1}, 100})

	cam := NewCamera()

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			var col Vector
			// Antialiasing. For each pixel, shoot <ns> random rays and average the color based on the hit.
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.ray(u, v)
				col = col.Add(pixel(world, r))
			}
			col = col.Scale(1.0 / float64(ns))

			ir := uint8(255.99 * col.R())
			ig := uint8(255.99 * col.G())
			ib := uint8(255.99 * col.B())

			img.Set(i, ny-j, color.RGBA{ir, ig, ib, 255})
		}
	}
	file, err := os.Create("demo.png")
	must(err)
	must(png.Encode(file, img))
}

func pixel(w World, r Ray) Vector {
	rec, hit := w.Hit(r, 0.0, math.MaxFloat64)
	if hit {
		return Vector{
			rec.Normal.X() + 1.0,
			rec.Normal.Y() + 1.0,
			rec.Normal.Z() + 1.0,
		}.Scale(0.5)
	}

	unitDirection := Unit(r.Direction())
	t := 0.5 * (unitDirection.Y() + 1.0)
	v1 := Vector{1.0, 1.0, 1.0}.Scale(1.0 - t)
	v2 := Vector{0.5, 0.7, 1.0}.Scale(t)
	return v1.Add(v2)
}
