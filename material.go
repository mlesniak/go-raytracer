package main

type Material interface {
	Scatter(r Ray, rec Hit) (scatter Ray, attenuation Vector, reflected bool)
}

type Lambertian struct {
	Albedo Vector
}

func (l Lambertian) Scatter(r Ray, rec Hit) (scatter Ray, attenuation Vector, reflected bool) {
	target := rec.P.Add(rec.Normal).Add(RandomInUnitSphere())
	scatter = Ray{rec.P, target.Sub(rec.P)}
	attenuation = l.Albedo
	return scatter, attenuation, true
}
