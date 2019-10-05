GHComp
======

#### A utility to compress and decompress sets of geohashes.

## Deflate

The Deflate utility accepts io.Reader containing newline-delimited geohashes of equal length, for example:

```
bdvkhunfnc90
bdvkj7vyvtz5
bdvkjkjbpr1t
bdvkjkn00jdn
bdvkjkjbremh
```

Duplicate geohashes will be discarded. The deflated output will contain only unique points.

The following example reads an input file, `data.txt`, and writes the deflated dataset to `deflated.txt`.

```go
package main

import (
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

## Inflate

The Inflate utility accepts a reader providing deflated geohash data as produced by the Deflate utility.

The following example reads an input file, `deflated.txt`, and writes the inflated dataset to `inflated.txt`.

```go
package main

import (
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
