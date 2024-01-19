# Geohash

This is a simple and thread-safe geohash implemented by [Go](https://go.dev/).

## Description

### Base32 Encoder
| Index | Binary Value | Base32 Character |
|-------|--------------|------------------|
| 0     | 00000        | 0                |
| 1     | 00001        | 1                |
| 2     | 00010        | 2                |
| 3     | 00011        | 3                |
| 4     | 00100        | 4                |
| 5     | 00101        | 5                |
| 6     | 00110        | 6                |
| 7     | 00111        | 7                |
| 8     | 01000        | 8                |
| 9     | 01001        | 9                |
| 10    | 01010        | B                |
| 11    | 01011        | C                |
| 12    | 01100        | D                |
| 13    | 01101        | E                |
| 14    | 01110        | F                |
| 15    | 01111        | G                |
| 16    | 10000        | H                |
| 17    | 10001        | J                |
| 18    | 10010        | K                |
| 19    | 10011        | M                |
| 20    | 10100        | N                |
| 21    | 10101        | P                |
| 22    | 10110        | Q                |
| 23    | 10111        | R                |
| 24    | 11000        | S                |
| 25    | 11001        | T                |
| 26    | 11010        | U                |
| 27    | 11011        | V                |
| 28    | 11100        | W                |
| 29    | 11101        | X                |
| 30    | 11110        | Y                |
| 31    | 11111        | Z                |


### Digits and precision
| Geohash Length | Lat bits | Lng bits | Lat err   | Lng err  | Err       |
|----------------|----------|----------|-----------|----------|-----------|
| 1              | 2        | 3        | ±23       | ±23      | ±2,500 km |
| 2              | 5        | 5        | ±2.8      | ±5.6     | ±630 km   |
| 3              | 7        | 8        | ±0.70     | ±0.70    | ±78 km    |
| 4              | 10       | 10       | ±0.087    | ±0.18    | ±20 km    |
| 5              | 12       | 13       | ±0.022    | ±0.022   | ±2.4 km   |
| 6              | 15       | 15       | ±0.0027   | ±0.0055  | ±610 m    |
| 7              | 17       | 18       | ±0.00068  | ±0.00068 | ±76 m     |
| 8              | 20       | 20       | ±0.000085 | ±0.00017 | ±19 m     |


### Width and Height of rectangle in latitude/longitude space
| Geohash length | Width     | Height    |
|----------------|-----------|-----------|
| 1              | 5009.4 km | 4992.6 km |
| 2              | 1252.3 km | 624.1 km  |
| 3              | 156.5 km  | 156 km    |
| 4              | 39.1 km   | 19.5 km   |
| 5              | 4.9 km    | 4.9 km    |
| 6              | 1.2 km    | 609.4 m   |
| 7              | 152.9 m   | 152.4 m   |
| 8              | 38.2 m    | 19 m      |


## Getting started

### Prerequisites
- **[Go](https://go.dev/) version 1.18+**

### Getting
With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```
import "github.com/ALong1997/geohash"
```

Otherwise, run the following Go command to install the `geohash` package:

```sh
$ go get -u https://github.com/ALong1997/geohash
```

### Quick Start

```go
package main

import (
    "fmt"

	"github.com/ALong1997/geohash"
)

func main() {
    t := geohash.NewTrie()

	p1 := geohash.NewPoint(13.361389, 38.115556, "Palermo")
	p2 := geohash.NewPoint(121.506377, 31.245105, "东方明珠")
	t.Put(p1)
	t.Put(p2)
	
	// SQC8B49R WTW3SZYP
    fmt.Println(p1.Geohash(), p2.Geohash())
	
    t.Get(geohash.Geohash("WTW3SZYP"))
    t.GetByPrefix("WTW")
	
    points, err :=t.GetPointsByCircle(p1, 10)
	if err != nil {
        fmt.Println(err)
	}

	for _, point := range points {
		fmt.Println(point.GetLng(), point.GetLat(), point.GetVal().(string))
	}
}

```
