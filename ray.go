package main

type Ray struct {
	// a + t*b
	a, b Vector
}

func (r *Ray) Origin() Vector {
	return r.a
}

func (r *Ray) Direction() Vector {
	return r.b
}

func (r *Ray) At(t float64) Vector {
	return r.a.Add(r.b.Scale(t))
}

type Hit struct {
	T      float64
	P      Vector
	Normal Vector
}

type Hiter interface {
	Hit(r Ray, tMin, tMax float64) (data Hit, hit bool)
}
