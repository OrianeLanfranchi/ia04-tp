package sprite

type Rectangle struct {
	topLeft     float64
	bottomRight float64
}

func (r Rectangle) TopLeft() float64 {
	return r.topLeft
}

func (r *Rectangle) SetTopLeft(tl float64) {
	r.topLeft = tl
}

func (r Rectangle) BottomRight() float64 {
	return r.bottomRight
}

func (r *Rectangle) SetBottomRight(br float64) {
	r.bottomRight = br
}
