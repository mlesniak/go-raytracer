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
	step := 1
	fmt.Printf("Computing %d pixel with aliasing=%d; == %dM pixels\n", (nx*ny)/step, ns, (nx*ny)/step*ns/1_000_000)

	world := World{}
	world.Add(Sphere{
		Vector{0, -1000.5, 0}, 1000,
		Lambertian{Albedo: Vector{0.5, 0.5, 0.5}}})
	world.Add(Sphere{
		Vector{0, 0, -1}, 0.5,
		Lambertian{Albedo: Vector{0.1, 0.2, 0.5}}})
	world.Add(Sphere{
		Vector{1, 0, -1}, 0.5,
		Metal{Albedo: Vector{0.8, 0.6, 0.2}, Fuzziness: 0.5}})
	world.Add(Sphere{
		Vector{-1, 0, -1}, 0.5,
		Dielectric{1.5}})
	world.Add(Sphere{
		Vector{-1, 0, -1}, -0.45,
		Dielectric{1.5}})

	lookFrom := Vector{3, 1, 2}
	lookAt := Vector{0, 0, -1}
	distToFocus := lookFrom.Sub(lookAt).Len()
	aperture := 2.0
	cam := NewCamera(
		lookFrom,
		lookAt,
		Up(),
		40, float64(nx)/float64(ny),
		aperture, distToFocus)

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))
	for j := ny - 1; j >= 0; j -= step {
		for i := 0; i < nx; i += step {
			var col Vector
			// Antialiasing. For each pixel, shoot <ns> random rays and average the color based on the hit.
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := cam.ray(u, v)
				col = col.Add(pixel(world, r, 0))
			}
			col = col.Scale(1.0 / float64(ns))
			col = Vector{math.Sqrt(col.R()), math.Sqrt(col.G()), math.Sqrt(col.B())}

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

func pixel(w World, r Ray, depth int) Vector {
	rec, hit := w.Hit(r, 0.001, math.MaxFloat64)
	if hit {
		scatter, attenuation, reflection := rec.Material.Scatter(r, rec)
		if depth < 50 && reflection {
			return attenuation.Mul(pixel(w, scatter, depth+1))
		} else {
			return Vector{0, 0, 0}
		}
	}

	unitDirection := Unit(r.Direction())
	t := 0.5 * (unitDirection.Y() + 1.0)
	v1 := Vector{1.0, 1.0, 1.0}.Scale(1.0 - t)
	v2 := Vector{0.5, 0.7, 1.0}.Scale(t)
	return v1.Add(v2)
}
