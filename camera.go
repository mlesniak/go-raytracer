package main

import "math"

type Camera struct {
	Origin     Vector
	LowerLeft  Vector
	Horizontal Vector
	Vertical   Vector
}

func NewCamera(lookFrom, lookAt, vup Vector, vfov float64, aspect float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight

	w := lookFrom.Sub(lookAt).Unit()
	u := vup.Cross(w).Unit()
	v := w.Cross(u).Unit()

	origin := lookFrom
	lowerLeft := origin.Sub(u.Scale(halfWidth)).Sub(v.Scale(halfHeight)).Sub(w)
	horizontal := u.Scale(2 * halfWidth)
	vertical := v.Scale(2 * halfHeight)

	return Camera{origin, lowerLeft, horizontal, vertical}
}

func (c *Camera) ray(s, t float64) Ray {
	return Ray{c.Origin,
		(c.LowerLeft.Add(c.Horizontal.Scale(s)).Add(c.Vertical.Scale(t))).Sub(c.Origin)}
}

func Up() Vector {
	return Vector{0, 1, 0}
}
