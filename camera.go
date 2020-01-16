package main

import "math"

type Camera struct {
	Origin     Vector
	LowerLeft  Vector
	Horizontal Vector
	Vertical   Vector

	u, v, w    Vector
	lensRadius float64
}

func NewCamera(lookFrom, lookAt, vup Vector, vfov float64, aspect, aperture, focusDist float64) (c Camera) {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight
	c.lensRadius = aperture / 2.0

	c.w = lookFrom.Sub(lookAt).Unit()
	c.u = vup.Cross(c.w).Unit()
	c.v = c.w.Cross(c.u).Unit()

	c.Origin = lookFrom
	c.LowerLeft = c.Origin.Sub(c.u.Scale(halfWidth * focusDist)).Sub(c.v.Scale(halfHeight * focusDist)).Sub(c.w.Scale(focusDist))
	c.Horizontal = c.u.Scale(2 * halfWidth * focusDist)
	c.Vertical = c.v.Scale(2 * halfHeight * focusDist)

	//return Camera{origin, lowerLeft, horizontal, vertical}
	return
}

func (c *Camera) ray(s, t float64) Ray {
	rd := RandomInUnitSphere().Scale(c.lensRadius)
	offset := c.u.Scale(rd.X()).Add(c.v.Scale(rd.Y()))

	return Ray{c.Origin.Add(offset),
		(c.LowerLeft.Add(c.Horizontal.Scale(s)).Add(c.Vertical.Scale(t))).Sub(c.Origin).Sub(offset)}
}

func Up() Vector {
	return Vector{0, 1, 0}
}
