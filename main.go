package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	nx := 200
	ny := 100
	ns := 1
	cores := -1
	step := 1

	rand.Seed(time.Now().UnixNano())

	world := World{}
	world.Add(Sphere{
		Vector{0, -1000.0, 0}, 1000,
		Lambertian{Albedo: Vector{0.1, 0.1, 0.1}}})

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			mat := rand.Float64()
			radius := rand.Float64()*(0.2-0.1) + 0.1
			center := Vector{float64(a) + 0.9*rand.Float64(), radius, float64(b) + 0.9*rand.Float64()}
			if center.Sub(Vector{0, 1, 0}).Len() > 1.5 {
				if mat < 0.5 {
					world.Add(Sphere{center, radius,
						Lambertian{
							Albedo: Vector{
								rand.Float64() * rand.Float64(),
								rand.Float64() * rand.Float64(),
								rand.Float64() * rand.Float64(),
							}}})
				} else if mat < 0.95 {
					world.Add(Sphere{center, radius,
						Metal{
							Albedo: Vector{
								rand.Float64(),
								rand.Float64(),
								rand.Float64(),
							},
							Fuzziness: 0.5 * rand.Float64(),
						}})
				} else {
					world.Add(Sphere{center, 0.1, Dielectric{1.5}})
				}
			}
		}
	}
	world.Add(Sphere{Vector{-2, 1, 0}, 1.0, Dielectric{1.5}})
	world.Add(Sphere{Vector{0.5, 1.5, 0}, 1.5, Metal{Vector{.9, .1, .1}, 0.3}})
	world.Add(Sphere{Vector{4, 2, 0}, 2.0, Metal{Vector{.7, .6, .5}, 0.0}})

	fmt.Printf("Computing %d pixel with aliasing=%d; == %dM pixels / %d objects / cores=%d\n", (nx*ny)/step, ns, (nx*ny)/step*ns/1_000_000, len(world.Objects), cores)

	lookFrom := Vector{2, 1.0, 8}
	lookAt := Vector{0, 1, 0}

	// Animation 0000--------------------------------------------------------------------------------------------------
	segments := 360
	angle := 360.0 / float64(segments)
	deg := angle

	for i := 1; i <= segments; i++ {
		fmt.Println("Rendering segment", i)
		distToFocus := lookFrom.Sub(lookAt).Len()
		aperture := 0.0
		cam := NewCamera(
			lookFrom,
			lookAt,
			Up(),
			40, float64(nx)/float64(ny),
			aperture, distToFocus)

		img := computeImage(cores, nx, ny, step, ns, cam, world)
		file, err := os.Create(fmt.Sprintf("demo-%02d.png", i))
		must(err)
		must(png.Encode(file, img))

		rad := deg * math.Pi / 180

		newX := math.Cos(rad)*(lookFrom.X()-lookAt.X()) - math.Sin(rad)*(lookFrom.Z()-lookAt.Z()) + lookAt.X()
		newY := math.Sin(rad)*(lookFrom.X()-lookAt.X()) + math.Cos(rad)*(lookFrom.Z()-lookAt.Z()) + lookAt.Z()
		lookFrom = Vector{newX, 1.0, newY}
	}

	// Single image --------------------------------------------------------------------------------------------------
	//distToFocus := lookFrom.Sub(lookAt).Len()
	//aperture := 0.0
	//cam := NewCamera(
	//	lookFrom,
	//	lookAt,
	//	Up(),
	//	40, float64(nx)/float64(ny),
	//	aperture, distToFocus)
	//
	//img := computeImage(cores, nx, ny, step, ns, cam, world)
	//file, err := os.Create("demo.png")
	//must(err)
	//must(png.Encode(file, img))

	duration := time.Now().Sub(start)
	fmt.Printf("Rendering took %10.4gs\n", duration.Seconds())
}

type result struct {
	row  int
	data []color.RGBA
}

func computeImage(cores int, nx int, ny int, step int, ns int, cam Camera, world World) *image.RGBA {
	results := make(chan result, ny)
	status := make(chan string)

	var sem Semaphore
	if cores <= 0 {
		sem = NewSemaphore(ny)
	} else {
		sem = NewSemaphore(cores)
	}

	img := image.NewRGBA(image.Rect(0, 0, nx, ny))
	var wg sync.WaitGroup
	wg.Add(ny)
	for j := ny - 1; j >= 0; j -= step {
		go func(j int) {
			sem.Acquire(1)
			defer sem.Release(1)
			//fmt.Println("Spawning row", j)
			row := make([]color.RGBA, nx)
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
				row[i] = color.RGBA{ir, ig, ib, 255}
			}
			results <- result{j, row}
			status <- fmt.Sprintf("Row %d finished", j)
		}(j)
	}
	finished := 0.0
	for i := 0; i < ny; i++ {
		<-status
		finished++
		perc := finished / float64(ny) * 100.0
		fmt.Printf("%3.0f%%        \r", perc)
	}
	//wg.Wait()
	close(results)

	for result := range results {
		for i, rgba := range result.data {
			img.Set(i, ny-result.row, rgba)
		}
	}

	return img
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
