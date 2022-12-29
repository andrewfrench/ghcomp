package ghcomp

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

var (
	ErrPrecision = errors.New("input value does not match configured precision")
)

const (
	GlobalStart = '!'
	SegmentStop = '.'
)

// Tree is the receiver for publicly exported methods.
type Tree struct {
	root      *node
	precision int
}

// New returns an empty, initialized tree.
func New(precision int) *Tree {
	return &Tree{
		precision: precision,
		root: &node{
			value:    GlobalStart,
			children: make(map[byte]*node),
		},
	}
}

// WriteInflated writes the inflated geohash values contained within the tree to an io.Writer.
func (t *Tree) WriteInflated(out io.Writer) error {
	return (&inflate{out: out, precision: t.precision}).descend(t.root)
}

// WriteDeflated writes the deflated geohash values contained within the tree to an io.Writer.
func (t *Tree) WriteDeflated(out io.Writer) error {
	return (&deflate{out: out}).descend(t.root)
}

// Entree adds a value to the tree. The precision (length) of the input value must match the tree configuration.
func (t *Tree) Entree(value []byte) error {
	if len(value) != t.precision {
		return ErrPrecision
	}

	t.root.entree(value)

	return nil
}

// EntreeDeflated adds deflated geohash values to the tree.
func (t *Tree) EntreeDeflated(in io.Reader) error {
	scanner := bufio.NewScanner(in)
	scanner.Split(ScanSegment)
	if !scanner.Scan() {
		return fmt.Errorf("failed to scan input source")
	}

	window := scanner.Bytes()
	err := t.Entree(window)
	if err != nil {
		return err
	}

	for scanner.Scan() {
		mask := scanner.Bytes()
		offset := len(window) - len(mask)
		for i := range mask {
			window[offset+i] = mask[i]
		}

		err = t.Entree(window)
		if err != nil {
			return err
		}
	}

	return nil
}

// ScanSegment reads a chunk of deflated geohash data delimited by SegmentStop.
var ScanSegment bufio.SplitFunc = func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF {
		return
	}

	advance = 0
	token = make([]byte, 0)
	for i := range data {
		advance++
		if data[i] == SegmentStop {
			return
		}

		token = append(token, data[i])
	}

	return
}
