package ghcomp

import (
	"bufio"
	"io"
)

// Deflate reads inflated geohash data from an io.Reader and writes deflated data to an io.Writer.
func Deflate(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	root := &node{
		value:    GlobalStart,
		children: make(map[byte]*node),
	}

	var geohash []byte
	for scanner.Scan() {
		geohash = scanner.Bytes()
		root.entree(geohash)
	}

	s := deflate{out: out}
	err := s.descend(root)
	if err != nil {
		return err
	}

	return nil
}

type deflate struct {
	prev byte
	out  io.Writer
}

func (d *deflate) descend(current *node) error {
	err := d.put(current.value)
	if err != nil {
		return err
	}

	for k := range current.children {
		err = d.descend(current.children[k])
		if err != nil {
			return err
		}
	}

	err = d.put(SegmentStop)
	if err != nil {
		return err
	}

	return nil
}

func (d *deflate) put(b byte) error {
	if b == GlobalStart {
		return nil
	}

	if b == SegmentStop && d.prev == SegmentStop {
		return nil
	}

	d.prev = b
	_, err := d.out.Write([]byte{b})
	if err != nil {
		return err
	}

	return nil
}
