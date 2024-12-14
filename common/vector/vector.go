package vector

type Vec struct {
	X, Y int
}

func New(x,y int) Vec{
	return Vec{x,y}
}

func (v Vec) Mag() int {
	return v.X +v.Y
}

func (v Vec) Scale(s int) Vec {
	return Vec{v.X*s, v.Y*s}
}

func (v Vec) Add(b Vec) Vec {
	return Vec{v.X + b.X, v.Y + b.Y}
}

func (v Vec) Equals(b Vec) bool {
	return v.X == b.X && v.Y == b.Y
}