package main

import (
	"log"
	"os"

	"github.com/andrewfrench/ghcomp/pkg/ghcomp"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Requires two inputs: an source file and a destination file")
	}

	in, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Failed to open source file: %v", err)
	}

	out, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatalf("Failed to create destination file: %v", err)
	}

	err = ghcomp.NewInflater(in, out).Inflate()
	if err != nil {
		log.Fatalf("Failed to inflate: %v", err)
	}
}
