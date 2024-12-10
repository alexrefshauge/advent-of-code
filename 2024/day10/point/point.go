package point

type Point struct {
	X, Y int
}

func NewP(x,y int) Point {
	return Point{X: x, Y: y}
}

func (p Point) Shift(x,y int) Point { 
	return Point{p.X + x, p.Y + y}
}