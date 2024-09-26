package sprite

type Sprite struct {
	position Point2D
	hitbox   Rectangle
	zoom     float64
	bitmap   string
}

func NewSprite(pos Point2D, rec Rectangle) *Sprite {
	s := new(Sprite)
	s.position = pos
	s.hitbox = rec
	s.zoom = 1
	s.bitmap = "NeverGonnaGiveYouUp"
	return s
}

func (s Sprite) Position() Point2D {
	return s.position
}

func (s *Sprite) SetPosition(p Point2D) {
	s.position = p
}

func (s Sprite) Hitbox() Rectangle {
	return s.hitbox
}

func (s *Sprite) SetHitbox(h Rectangle) {
	s.hitbox = h
}

func (s Sprite) Zoom() float64 {
	return s.zoom
}

func (s *Sprite) SetZoom(z float64) {
	s.zoom = z
}

func (s Sprite) Bitmap() string {
	return s.bitmap
}

func (s *Sprite) SetBitmap(bm string) {
	s.bitmap = bm
}

func (s *Sprite) Move() {
	//todo
}

/*func (s1 Sprite) Collision(s2 Sprite) Rectangle {

}*/
