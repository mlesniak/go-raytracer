package main

type Sphere struct {
	Center Vector
	Radius float64
}

func (s Sphere) Hit(r Ray) bool {
	oc := r.Origin().Sub(s.Center)
	a := Dot(r.Direction(), r.Direction())
	b := 2.0 * Dot(oc, r.Direction())
	c := Dot(oc, oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}
