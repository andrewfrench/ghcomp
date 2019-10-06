GHComp
======

#### A utility to compress and decompress sets of geohashes.

The inflate and deflate utilities do not preserve duplicate geohashes or the order of the dataset.

## Deflate

The Deflate utility accepts an `io.Reader` containing newline-delimited geohashes of equal length, for example:

```
...
bdvkhunfnc90
bdvkj7vyvtz5
bdvkjkjbpr1t
bdvkjkn00jdn
bdvkjkjbremh
...
```

The following example reads an input file, `data.txt`, and writes the deflated dataset to `deflated.txt`.

```go
package main

import (
	"os"
	"log"
	
	"github.com/andrewfrench/ghcomp"
)

func main() {
	in, _ := os.Open("data.txt")
	out, _ := os.Create("deflated.txt")

	err := ghcomp.NewDeflater(in, out).Deflate()
	if err != nil {
		log.Fatalf("failed to deflate data: %v", err)
	}
}
```

Deflate geohash data directly from the terminal using `cmd/deflate/deflate.go`:

```bash
$ go run cmd/deflate/deflate.go data.txt deflated.txt
```

## Inflate

The Inflate utility accepts an `io.Reader` providing deflated geohash data as produced by the Deflate utility.

The following example reads an input file, `deflated.txt`, and writes the inflated dataset to `inflated.txt`.

```go
package main

import (
	"os"
	"log"
	
	"github.com/andrewfrench/ghcomp"
)

func main() {
	in, _ := os.Open("deflated.txt")
	out, _ := os.Create("inflated.txt")

	err := ghcomp.NewInflater(in, out).Inflate()
	if err != nil {
		log.Fatalf("failed to inflate data: %v", err)
	}
}
```

Inflate geohash data directly from the terminal using `cmd/inflate/inflate.go`:

```bash
$ go run cmd/inflate/inflate.go deflated.txt inflated.txt
```