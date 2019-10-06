package ghcomp

import (
	"bufio"
	"fmt"
	"io"
)

type Inflater interface {
	Inflate() error
}

type inflater struct {
	in  io.Reader
	out io.Writer
}

func NewInflater(in io.Reader, out io.Writer) Inflater {
	d := &inflater{
		in:  in,
		out: out,
	}

	return d
}

func (i *inflater) Inflate() error {
	scanner := bufio.NewScanner(i.in)
	scanner.Split(ScanSegment)
	if !scanner.Scan() {
		return fmt.Errorf("failed to scan input source")
	}

	window := scanner.Bytes()
	_, err := i.out.Write(append(window, '\n'))
	if err != nil {
		return err
	}

	for scanner.Scan() {
		mask := scanner.Bytes()
		offset := len(window) - len(mask)
		for i := range mask {
			window[offset+i] = mask[i]
		}

		_, err := i.out.Write(append(window, '\n'))
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
