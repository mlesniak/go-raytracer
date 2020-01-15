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
