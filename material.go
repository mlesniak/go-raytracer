package main

import (
	"math"
	"math/rand"
)

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

type Metal struct {
	Albedo    Vector
	Fuzziness float64
}

func (m Metal) Scatter(r Ray, rec Hit) (scatter Ray, attenuation Vector, reflected bool) {
	ref := Reflect(Unit(r.Direction()), rec.Normal)
	scatter = Ray{rec.P, ref}
	fuzz := 1.0
	if m.Fuzziness < 1.0 {
		fuzz = m.Fuzziness
	}
	scatter = Ray{rec.P, ref.Add(RandomInUnitSphere().Scale(fuzz))}
	attenuation = m.Albedo
	return scatter, attenuation, Dot(scatter.Direction(), rec.Normal) > 0
}

func Reflect(v, n Vector) Vector {
	d := 2.0 * Dot(v, n)
	n2 := n.Scale(d)
	return v.Sub(n2)
}

func Schlick(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

func Refract(v, n Vector, niOverNt float64) (ref Vector, refracted bool) {
	uv := Unit(v)
	dt := Dot(uv, n)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		p1 := uv.Sub(n.Scale(dt)).Scale(niOverNt)
		p2 := n.Scale(math.Sqrt(discriminant))
		return p1.Sub(p2), true
	}

	return ref, false
}

type Dielectric struct {
	refractionIndex float64
}

func (d Dielectric) Scatter(r Ray, rec Hit) (scatter Ray, attenuation Vector, reflected bool) {
	var outwardNormal Vector
	var niOverNt float64

	var reflectProb float64
	var cosine float64

	if Dot(r.Direction(), rec.Normal) > 0 {
		outwardNormal = rec.Normal.Scale(-1.0)
		niOverNt = d.refractionIndex
		cosine = d.refractionIndex * Dot(r.Direction(), rec.Normal) / r.Direction().Len()
	} else {
		outwardNormal = rec.Normal
		niOverNt = 1.0 / d.refractionIndex
		cosine = -Dot(r.Direction(), rec.Normal) / r.Direction().Len()
	}

	attenuation = Vector{1.0, 1.0, 0.0}
	refracted, ok := Refract(r.Direction(), outwardNormal, niOverNt)
	if ok {
		reflectProb = Schlick(cosine, d.refractionIndex)
		//scatter = Ray{rec.P, refracted}
		//return scatter, attenuation, true
	} else {
		reflectProb = 1.0
		//refl := Reflect(r.Direction(), rec.Normal)
		//scatter = Ray{rec.P, refl}
		//return scatter, attenuation, false
	}

	if rand.Float64() < reflectProb {
		refl := Reflect(r.Direction(), rec.Normal)
		scatter = Ray{rec.P, refl}
		return scatter, attenuation, true
	} else {
		scatter = Ray{rec.P, refracted}
		return scatter, attenuation, true
	}
}
