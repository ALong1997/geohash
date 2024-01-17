package geohash

import (
	"errors"
	"strconv"
	"strings"
)

type Geohash string

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

var (
	encoder = [32]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J', 'K', 'M', 'N',
		'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

	decoder = map[byte]uint8{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9,
		'B': 10, 'C': 11, 'D': 12, 'E': 13, 'F': 14, 'G': 15, 'H': 16, 'J': 17, 'K': 18, 'M': 19, 'N': 20,
		'P': 21, 'Q': 22, 'R': 23, 'S': 24, 'T': 25, 'U': 26, 'V': 27, 'W': 28, 'X': 29, 'Y': 30, 'Z': 31}

	/*
		geohash length | width | height
			1 | 5009.4km | 4992.6km
			2 | 1252.3km | 624.1km
			3 | 156.5km | 156km
			4 | 39.1km | 19.5km
			5 | 4.9km | 4.9km
			6 | 1.2km | 609.4m
			7 | 152.9m | 152.4m
			8 | 38.2m | 19m
			9 | 4.8m | 4.8m
			10 | 1.2m | 59.5cm
			11 | 14.9cm | 14.9m
			12 | 3.7cm | 1.9cm
	*/
	maxIncircleDiameterRank = [geohashLen]uint32{4992600, 624100, 156000, 19500, 4900, 609, 152, 19}
)

var ErrInvalidDiameter = errors.New("invalid diameter")

// GetGeohash converts the longitude and latitude into corresponding fixed 40-bit geohash strings,
// 5 bits is mapped by one base32, so it consists of a total of 8 base32 characters.
func GetGeohash(point *Point) Geohash {
	if point == nil {
		return ""
	}
	geohash := strings.Builder{}
	lngBits := encode(point.GetLng(), minLng, maxLng)
	latBits := encode(point.GetLat(), minLat, maxLat)

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

// check whether geohash is valid
func (g Geohash) check() bool {
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
func getGeohashLenByDiameter(diameter uint32) (uint8, error) {
	if diameter == 0 || diameter > maxIncircleDiameterRank[geohashLen-1] {
		return 0, ErrInvalidDiameter
	}

	for i := geohashLen - 1; i >= 0; i-- {
		if maxIncircleDiameterRank[i] >= diameter {
			return uint8(i), nil
		}
	}

	return 0, ErrInvalidDiameter
}
