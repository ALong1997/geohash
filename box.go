package geohash

// Box is a rectangle in latitude/longitude space
type Box struct {
	Geohash Geohash
	Points  map[string]*Point
}

func NewBox(geohash Geohash, points map[string]*Point) *Box {
	if !geohash.check() {
		return nil
	}
	return &Box{
		Geohash: geohash,
		Points:  points,
	}
}

func (b *Box) GetGeohash() Geohash {
	if b == nil {
		return ""
	}
	return b.Geohash
}

func (b *Box) GetPoints() map[string]*Point {
	if b == nil {
		return nil
	}
	return b.Points
}

func (b *Box) GetAllPoints() []*Point {
	if b == nil || len(b.Points) == 0 {
		return []*Point{}
	}

	res := make([]*Point, 0, len(b.Points))
	for _, point := range b.Points {
		res = append(res, point)
	}
	return res
}

func (b *Box) add(point *Point) {
	if b == nil || !b.Geohash.check() || point == nil {
		return
	}

	if len(b.Points) == 0 {
		b.Points = map[string]*Point{}
	}
	b.Points[point.key()] = point
}
