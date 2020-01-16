package main

type Camera struct {
	Origin     Vector
	LowerLeft  Vector
	Horizontal Vector
	Vertical   Vector
}

func NewCamera() Camera {
	origin := Vector{0.0, 0.0, 0.0}
	lowerLeft := Vector{-2.0, -1.0, -1.0}
	horizontal := Vector{4.0, 0.0, 0.0}
	vertical := Vector{0.0, 2.0, 0.0}

	return Camera{origin, lowerLeft, horizontal, vertical}
}

func (c *Camera) ray(u, v float64) Ray {
	return Ray{c.Origin, c.LowerLeft.Add(c.Horizontal.Scale(u).Add(c.Vertical.Scale(v)))}
}
