package geohash

import "fmt"

type Point struct {
	Lng, Lat float64
}

func NewPoint(lng, lat float64) *Point {
	return &Point{
		Lng: lng,
		Lat: lat,
	}
}

func (p *Point) GetLng() float64 {
	if p == nil {
		return 0
	}
	return p.Lng
}
func (p *Point) GetLat() float64 {
	if p == nil {
		return 0
	}
	return p.Lat
}

func (p *Point) key() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("%v_%v", p.Lng, p.Lat)
}
