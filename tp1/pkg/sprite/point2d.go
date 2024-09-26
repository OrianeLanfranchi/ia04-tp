package sprite

import "math"

type Point2D struct {
	x float64
	y float64
}

func NewPoint2D(x, y float64) *Point2D {
	p := new(Point2D)
	p.x = x
	p.y = y
	return p
}

func (p Point2D) X() float64 {
	return p.x
}

func (p *Point2D) SetX(x float64) {
	p.x = x
}

func (p Point2D) Y() float64 {
	return p.y
}

func (p *Point2D) SetY(y float64) {
	p.y = y
}

func (p Point2D) Clone() Point2D {
	return p
}

func (p Point2D) Module() float64 {
	return math.Sqrt(math.Pow(0-p.x, 2) + math.Pow(0-p.y, 2))
}
