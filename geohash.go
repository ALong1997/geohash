package geohash

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	minLat = -90
	maxLat = 90
	minLng = -180
	maxLng = 180
)

const (
	bit0 = '0'
	bit1 = '1'

	bitsLen    = 20
	geohashLen = bitsLen << 1 / 5

	invalidCode = 32
)

const (
	earthRadius   = 6371000
	coordinateToM = 111000 // estimated value of latitude and longitude converted into meters
)

var (
	encoder = [32]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'M', 'N',
		'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	decoder = map[byte]uint8{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'B': 10, 'C': 11, 'D': 12, 'E': 13, 'F': 14, 'G': 15, 'H': 16, 'J': 17, 'K': 18, 'M': 19, 'N': 20,
		'P': 21, 'Q': 22, 'R': 23, 'S': 24, 'T': 25, 'U': 26, 'V': 27, 'W': 28, 'X': 29, 'Y': 30, 'Z': 31}

	incircleDiameterRank = [geohashLen]uint32{4992600, 624100, 156000, 19500, 4900, 609, 152, 19}
)

var ErrInvalidDiameter = errors.New("invalid diameter")

type Geohash string

func (g Geohash) valid() bool {
	if len(g) != geohashLen {
		return false
	}

	for i := 0; i < geohashLen; i++ {
		if decode(g[i]) == invalidCode {
			return false
		}
	}
	return true
}

type Point struct {
	Lng, Lat float64
	Val      any
}

func NewPoint(lng, lat float64, val any) *Point {
	return &Point{
		Lng: lng,
		Lat: lat,
		Val: val,
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

func (p *Point) GetVal() (val any) {
	if p == nil {
		return
	}
	return p.Val
}

func (p *Point) Distance(target *Point) uint32 {
	if p == nil || target == nil {
		return 0
	}
	return uint32(haversine(p, target))
}

// Geohash converts the longitude and latitude into corresponding fixed 40-bit geohash strings,
// 5 bits is mapped by one base32, so it consists of a total of 8 base32 characters.
func (p *Point) Geohash() Geohash {
	if p == nil {
		return ""
	}
	geohash := strings.Builder{}
	lngBits := encode(p.Lng, minLng, maxLng)
	latBits := encode(p.Lat, minLat, maxLat)

	mixBits := strings.Builder{}
	for i := 1; i <= bitsLen<<1; i++ {
		if i&1 == 1 {
			mixBits.WriteByte(lngBits[(i-1)>>1])
		} else {
			mixBits.WriteByte(latBits[(i-1)>>1])
		}

		if i%5 == 0 {
			i, _ := strconv.ParseUint(mixBits.String(), 2, 8)
			geohash.WriteByte(encoder[i])
			mixBits.Reset()
		}
	}

	return Geohash(geohash.String())
}

func (p *Point) key() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("%v_%v", p.Lng, p.Lat)
}

// circumscribedSquarePointsByCircle return the circumscribed square of the circle with point and radius
func (p *Point) circumscribedSquarePointsByCircle(radius uint32) [9]*Point {
	dif := float64(radius) / coordinateToM
	left := p.Lng - dif
	if left < minLng {
		left += maxLng << 1
	}
	right := p.Lng + dif
	if right > maxLng {
		right -= maxLng << 1
	}
	bot := p.Lat - dif
	if bot < minLat {
		bot += maxLat << 1
	}
	top := p.Lat + dif
	if top > maxLat {
		top -= maxLat << 1
	}

	return [9]*Point{
		NewPoint(left, top, nil),
		NewPoint(p.Lng, top, nil),
		NewPoint(right, top, nil),
		NewPoint(left, p.Lat, nil),
		NewPoint(p.Lng, p.Lat, nil),
		NewPoint(right, p.Lat, nil),
		NewPoint(left, bot, nil),
		NewPoint(p.Lng, bot, nil),
		NewPoint(right, bot, nil),
	}
}

// encode converts the latitude or longitude coordinate into corresponding fixed 20-bit binary string
func encode(coordinate, start, end float64) string {
	bits := strings.Builder{}
	for i := 0; i < bitsLen; i++ {
		mid := (start + end) / 2
		if coordinate < mid {
			bits.WriteByte(bit0)
			end = mid
		} else {
			bits.WriteByte(bit1)
			start = mid
		}
	}
	return bits.String()
}

// decode converts the bit into corresponding decimal uint8
func decode(bit byte) uint8 {
	val, ok := decoder[bit]
	if ok {
		return val
	}
	return invalidCode
}

// getGeohashLenByDiameter returns the minimum length of geohash required by the circumscribed rectangle,
// according to the diameter of the circle
// diameter <= incircleDiameterRank[2]
func getGeohashLenByDiameter(diameter uint32) (uint8, error) {
	if diameter == 0 || diameter > incircleDiameterRank[2] {
		return 0, ErrInvalidDiameter
	}

	for i := geohashLen - 1; i >= 0; i-- {
		if incircleDiameterRank[i] >= diameter {
			return uint8(i) + 1, nil
		}
	}

	return 0, ErrInvalidDiameter
}

// haversine formula is used to calculate the distance of large circle route between two latitude and longitude coordinates.
// a = sin²((lat₂ - lat₁)/2) + cos(lat₁) * cos(lat₂) * sin²((lng₂ - lng₁)/2)
// c = 2 * atan2(√a, √(1−a))
// d = R * c
func haversine(p1, p2 *Point) float64 {
	if p1 == nil || p2 == nil {
		return 0
	}

	// 将经纬度转换为弧度
	radianLat1 := p1.Lat * (math.Pi / 180)
	radianLat2 := p2.Lat * (math.Pi / 180)
	radianDifLat := (p2.Lat - p1.Lat) * (math.Pi / 180)
	radianDifLng := (p2.Lng - p1.Lng) * (math.Pi / 180)

	// 应用Haversine公式
	a := math.Sin(radianDifLat/2)*math.Sin(radianDifLat/2) + math.Cos(radianLat1)*math.Cos(radianLat2)*math.Sin(radianDifLng/2)*math.Sin(radianDifLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := earthRadius * c

	return d
}
