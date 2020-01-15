package main

import "math"

type Vector struct {
	e0, e1, e2 float64
}

func (v Vector) R() float64 {
	return v.e0
}

func (v Vector) G() float64 {
	return v.e1
}

func (v Vector) B() float64 {
	return v.e2
}

func (v Vector) X() float64 {
	return v.e0
}

func (v Vector) Y() float64 {
	return v.e1
}

func (v Vector) Z() float64 {
	return v.e2
}

func (v Vector) Len() float64 {
	return math.Sqrt(v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2)
}

func (v Vector) SquaredLen() float64 {
	return v.e0*v.e0 + v.e1*v.e1 + v.e2*v.e2
}

func Unit(v Vector) Vector {
	return v.Scale(1 / v.Len())
}

func (v Vector) Unit() {
	k := v.Len()
	v.e0 *= k
	v.e1 *= k
	v.e2 *= k
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		e0: v.e0 + v2.e0,
		e1: v.e1 + v2.e1,
		e2: v.e2 + v2.e2,
	}
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{
		e0: v.e0 - v2.e0,
		e1: v.e1 - v2.e1,
		e2: v.e2 - v2.e2,
	}
}

func (v Vector) Mul(v2 Vector) Vector {
	return Vector{
		e0: v.e0 * v2.e0,
		e1: v.e1 * v2.e1,
		e2: v.e2 * v2.e2,
	}
}

func (v Vector) Div(v2 Vector) Vector {
	return Vector{
		e0: v.e0 / v2.e0,
		e1: v.e1 / v2.e1,
		e2: v.e2 / v2.e2,
	}
}

func (v Vector) Scale(t float64) Vector {
	return Vector{
		e0: v.e0 * t,
		e1: v.e1 * t,
		e2: v.e2 * t,
	}
}

func (v Vector) Dot(v2 Vector) float64 {
	return v.e0*v2.e0 + v.e1*v2.e1 + v.e2*v2.e2
}

func (v Vector) Cross(v2 Vector) Vector {
	return Vector{
		e0: v.e1*v2.e2 - v.e2*v2.e1,
		e1: v.e0*v2.e2 - v.e2*v2.e0,
		e2: v.e0*v2.e1 - v.e1*v2.e0,
	}
}
