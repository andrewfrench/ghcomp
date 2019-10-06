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

type n struct {
	b byte
	m map[byte]*n
	p *n
}

func NewDeflater(in io.Reader, out io.Writer) Deflater {
	s := &deflater{
		in:  in,
		out: out,
	}

	return s
}

func (d *deflater) Deflate() error {
	scanner := bufio.NewScanner(d.in)
	scanner.Split(bufio.ScanLines)

	t := &n{
		b: GlobalStart,
		m: make(map[byte]*n),
	}

	var l []byte
	for scanner.Scan() {
		l = scanner.Bytes()
		d.entree(t, l)
	}

	err := d.descend(t)
	if err != nil {
		return err
	}

	return nil
}

func (d *deflater) entree(t *n, l []byte) {
	if len(l) == 0 {
		return
	}

	if t.m[l[0]] == nil {
		t.m[l[0]] = &n{
			b: l[0],
			m: make(map[byte]*n),
			p: t,
		}
	}

	d.entree(t.m[l[0]], l[1:])
}

func (d *deflater) descend(t *n) error {
	err := d.put(t.b)
	if err != nil {
		return err
	}

	for k := range t.m {
		err := d.descend(t.m[k])
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
