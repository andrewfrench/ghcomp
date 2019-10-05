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
	v bool
}

func NewDeflater(in io.Reader, out io.Writer) Deflater {
	s := &deflater{
		in:  in,
		out: out,
	}

	return s
}

func (s *deflater) Deflate() error {
	scanner := bufio.NewScanner(s.in)
	scanner.Split(bufio.ScanLines)

	t := &n{
		b: GlobalStart,
		m: make(map[byte]*n),
	}

	var l []byte
	for scanner.Scan() {
		l = scanner.Bytes()
		s.entree(t, l)
	}

	err := s.descend(t)
	if err != nil {
		return err
	}

	return nil
}

func (s *deflater) entree(t *n, l []byte) {
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

	s.entree(t.m[l[0]], l[1:])
}

func (s *deflater) descend(t *n) error {
	err := s.put(t.b)
	if err != nil {
		return err
	}

	for k := range t.m {
		if !t.m[k].v {
			err := s.descend(t.m[k])
			if err != nil {
				return err
			}
		}
	}

	err = s.put(SegmentStop)
	if err != nil {
		return err
	}

	return nil
}

func (s *deflater) put(b byte) error {
	if b == GlobalStart {
		return nil
	}

	if b == SegmentStop && s.prev == SegmentStop {
		return nil
	}

	s.prev = b
	_, err := s.out.Write([]byte{b})
	if err != nil {
		return err
	}

	return nil
}
