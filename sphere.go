package main

import "math"

type Sphere struct {
	Center Vector
	Radius float64
}

func (s Sphere) Hit(r Ray, tMin, tMax float64) (Hit, bool) {
	oc := r.Origin().Sub(s.Center)
	a := Dot(r.Direction(), r.Direction())
	b := 2.0 * Dot(oc, r.Direction())
	c := Dot(oc, oc) - s.Radius*s.Radius
	discriminant := b*b - 4*a*c

	var rec Hit
	if discriminant > 0 {
		tmp := (-b - math.Sqrt(discriminant)) / a
		if tmp < tMax && tmp > tMin {
			rec.T = tmp
			rec.P = r.At(rec.T)
			rec.Normal = rec.P.Sub(s.Center).Scale(1 / s.Radius)
			return rec, true
		}

		tmp = (-b + math.Sqrt(discriminant)) / a
		if tmp < tMax && tmp > tMin {
			rec.T = tmp
			rec.P = r.At(rec.T)
			rec.Normal = rec.P.Sub(s.Center).Scale(1 / s.Radius)
			return rec, true
		}
	}

	return rec, false
}
