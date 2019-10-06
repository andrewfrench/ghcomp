package ghcomp

import (
	"bufio"
	"io"
)

type Deflater interface {
	Deflate() error
}

type deflater struct {
	in   io.Reader
	out  io.Writer
	prev byte
}

type node struct {
	value    byte
	children map[byte]*node
	parent   *node
}

func NewDeflater(in io.Reader, out io.Writer) Deflater {
	d := &deflater{
		in:  in,
		out: out,
	}

	return d
}

func (d *deflater) Deflate() error {
	scanner := bufio.NewScanner(d.in)
	scanner.Split(bufio.ScanLines)

	root := &node{
		value:    GlobalStart,
		children: make(map[byte]*node),
	}

	var geohash []byte
	for scanner.Scan() {
		geohash = scanner.Bytes()
		d.entree(root, geohash)
	}

	err := d.descend(root)
	if err != nil {
		return err
	}

	return nil
}

func (d *deflater) entree(current *node, remaining []byte) {
	if len(remaining) == 0 {
		return
	}

	if current.children[remaining[0]] == nil {
		current.children[remaining[0]] = &node{
			value:    remaining[0],
			children: make(map[byte]*node),
			parent:   current,
		}
	}

	d.entree(current.children[remaining[0]], remaining[1:])
}

func (d *deflater) descend(current *node) error {
	err := d.put(current.value)
	if err != nil {
		return err
	}

	for k := range current.children {
		err := d.descend(current.children[k])
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

func (d *deflater) put(b byte) error {
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
