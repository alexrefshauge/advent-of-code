package main

type vec struct {
	x int
	y int
}

func (a vec) sub(b vec) vec {
	return vec{
		a.x - b.x,
		a.y - b.y,
	}
}

func (a vec) add(b vec) vec {
	return vec{
		a.x + b.x,
		a.y + b.y,
	}
}

func (a vec) scale(fac int) vec {
	return vec{
		a.x * fac,
		a.y * fac,
	}
}

func (v vec) inside() bool {
	return v.x >= 0 && v.y >= 0 && v.x < WIDTH && v.y < HEIGHT
}