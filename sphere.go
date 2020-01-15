package main

import "math"

type Sphere struct {
	Center Vector
	Radius float64
}

// Return 't' which allows to compute the point on which we collided.
func (s Sphere) Hit(r Ray) float64 {
	oc := r.Origin().Sub(s.Center)
	a := Dot(r.Direction(), r.Direction())
	b := 2.0 * Dot(oc, r.Direction())
	c := Dot(oc, oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return -1.0
	}

	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}
