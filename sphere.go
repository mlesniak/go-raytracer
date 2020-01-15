// TODO Define a type Sphere
package main

func HitSphere(center Vector, radius float64, r Ray) bool {
	oc := r.Origin().Sub(center)
	a := Dot(r.Direction(), r.Direction())
	b := 2.0 * Dot(oc, r.Direction())
	c := Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}
