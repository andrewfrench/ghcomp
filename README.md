GHComp
======

#### A utility to compress and decompress geohash sets.

The Deflate utility removes redundant geohash information from sets of geohashes. The ratio of deflated data to inflated data is variable and depends on the clustering of the dataset; those containing points that are physically closer to each other will share a larger portion of each geohash, allowing for smaller deflated datasets. In datasets containing points that are physically far from each other, each geohash will be have more unique data and cannot be as aggressively deflated. The Inflate and Deflate utilities do not preserve duplicate geohashes or the order of the dataset.

## Deflate

The Deflate utility accepts an `io.Reader` containing newline-delimited geohashes of equal length, for example:

```
bdvkhunfnc90
bdvkj7vyvtz5
bdvkjkjbpr1t
bdvkjkn00jdn
bdvkjkjbremh
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

	err := ghcomp.Deflate(in, out)
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

The Inflate utility accepts an `io.Reader` providing deflated geohash data as produced by the Deflate utility, for example:

```
bdvkhunfnc90.j7vyvtz5.kjbremh.pr1t.n00jdn.
```

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

	err := ghcomp.Inflate(in, out)
	if err != nil {
		log.Fatalf("failed to inflate data: %v", err)
	}
}
```

Inflate geohash data directly from the terminal using `cmd/inflate/inflate.go`:

```bash
$ go run cmd/inflate/inflate.go deflated.txt inflated.txt
```

The ghcomp tree can be interacted with directly as well. You can load deflated sets into the tree, add geohashes, combine sets, then write the tree in its inflated or deflated form.
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

	precision := 12
	tree := ghcomp.New(precision)
	err := tree.EntreeDeflated(in)
	if err != nil {
		log.Fatalf("failed to load deflated data: %v", err)
    }
	
	err = tree.Entree("abcabcabcabc")
	if err != nil {
		log.Fatalf("failed to entree: %v", err)
    }
	
	err = tree.WriteInflated(out)
	if err != nil {
		log.Fatalf("failed to write inflated data: %v", err)
	}
}
```