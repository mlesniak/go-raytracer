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

type Metal struct {
	Albedo Vector
}

func (m Metal) Scatter(r Ray, rec Hit) (scatter Ray, attenuation Vector, reflected bool) {
	ref := Reflect(Unit(r.Direction()), rec.Normal)
	scatter = Ray{rec.P, ref}
	attenuation = m.Albedo
	return scatter, attenuation, Dot(scatter.Direction(), rec.Normal) > 0
}

func Reflect(v, n Vector) Vector {
	d := 2.0 * Dot(v, n)
	n2 := n.Scale(d)
	return v.Sub(n2)
}
