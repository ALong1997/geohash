# Geohash

This is a simple and thread-safe geohash implemented by [Go](https://go.dev/).

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
