package ghcomp

import (
	"bufio"
	"fmt"
	"io"
)

type Deserializer interface {
	Deserialize() error
}

type deserializer struct {
	in  io.Reader
	out io.Writer
}

func NewDeserializer(in io.Reader, out io.Writer) Deserializer {
	d := &deserializer{
		in:  in,
		out: out,
	}

	return d
}

func (d *deserializer) Deserialize() error {
	s := bufio.NewScanner(d.in)
	s.Split(ScanSegment)
	if !s.Scan() {
		return fmt.Errorf("failed to scan input source")
	}

	w := s.Bytes()
	_, err := d.out.Write(w)
	if err != nil {
		return err
	}

	for s.Scan() {
		seg := s.Bytes()
		off := len(w) - len(seg)
		for i := range seg {
			w[off+i] = seg[i]
		}

		_, err := d.out.Write(w)
		if err != nil {
			return err
		}
	}

	return nil
}

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
