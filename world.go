package main

type World struct {
	Objects []Hiter
}

func (w *World) Hit(r Ray, tMin, tMax float64) (data Hit, hit bool) {
	var rec Hit
	hitAnything := false
	closestT := tMax

	for _, object := range w.Objects {
		trec, hit := object.Hit(r, tMin, closestT)
		if hit {
			hitAnything = true
			closestT = trec.T
			rec = trec
		}
	}

	return rec, hitAnything
}

func (w *World) Add(object Hiter) {
	w.Objects = append(w.Objects, object)
}
