package geohash

// Box is a rectangle in latitude/longitude space
type Box struct {
	Geohash  Geohash
	PointSet map[string]*Point
}

func NewBox(geohash Geohash, pointSet map[string]*Point) *Box {
	if !geohash.valid() {
		return nil
	}
	return &Box{
		Geohash:  geohash,
		PointSet: pointSet,
	}
}

func (b *Box) GetGeohash() Geohash {
	if b == nil {
		return ""
	}
	return b.Geohash
}

func (b *Box) GetPointSet() map[string]*Point {
	if b == nil {
		return nil
	}
	return b.PointSet
}

func (b *Box) GetAllPoints() []*Point {
	if b == nil || len(b.PointSet) == 0 {
		return []*Point{}
	}

	res := make([]*Point, 0, len(b.PointSet))
	for _, point := range b.PointSet {
		res = append(res, point)
	}
	return res
}

func (b *Box) add(point *Point) {
	if b == nil || !b.Geohash.valid() || point == nil {
		return
	}

	if len(b.PointSet) == 0 {
		b.PointSet = map[string]*Point{}
	}
	b.PointSet[point.key()] = point
}
