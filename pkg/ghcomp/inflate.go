package ghcomp

import (
	"bufio"
	"fmt"
	"io"
)

// Inflate reads deflated geohash data in an io.Reader and writes the inflated values to an io.Writer. This function
// can inflate the deflated values directly without using a tree.
func Inflate(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)
	scanner.Split(ScanSegment)
	if !scanner.Scan() {
		return fmt.Errorf("failed to scan input source")
	}

	window := scanner.Bytes()
	_, err := out.Write(append(window, '\n'))
	if err != nil {
		return err
	}

	for scanner.Scan() {
		mask := scanner.Bytes()
		offset := len(window) - len(mask)
		for i := range mask {
			window[offset+i] = mask[i]
		}

		_, err = out.Write(append(window, '\n'))
		if err != nil {
			return err
		}
	}

	return nil
}

type inflate struct {
	out       io.Writer
	precision int
}

func (in *inflate) descend(current *node) error {
	var err error
	for k := range current.children {
		err = in.descend(current.children[k])
		if err != nil {
			return err
		}
	}

	if len(current.children) == 0 {
		n := current
		val := make([]byte, in.precision)
		for i := 0; i < in.precision; i++ {
			if n.parent == nil {
				break
			}

			val[in.precision-i-1] = n.value
			n = n.parent
		}

		_, err = in.out.Write(append(val, '\n'))
		if err != nil {
			return err
		}
	}

	return nil
}
