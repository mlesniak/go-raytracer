package main

import "math"

type Camera struct {
	Origin     Vector
	LowerLeft  Vector
	Horizontal Vector
	Vertical   Vector
}

func NewCamera(vfov float64, aspect float64) Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight

	origin := Vector{0.0, 0.0, 0.0}
	lowerLeft := Vector{-halfWidth, -halfHeight, -1.0}
	horizontal := Vector{2 * halfWidth, 0.0, 0.0}
	vertical := Vector{0.0, 2 * halfHeight, 0.0}

	return Camera{origin, lowerLeft, horizontal, vertical}
}

func (c *Camera) ray(u, v float64) Ray {
	return Ray{c.Origin, c.LowerLeft.Add(c.Horizontal.Scale(u).Add(c.Vertical.Scale(v)))}
}
