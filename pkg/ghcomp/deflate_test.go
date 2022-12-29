package ghcomp

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDeflater_Deflate(t *testing.T) {
	in := bytes.NewBufferString(inflated)
	out := bytes.NewBuffer(nil)

	err := Deflate(in, out)
	assert.NoError(t, err)

	in = bytes.NewBuffer(out.Bytes())
	out = bytes.NewBuffer(nil)
	err = Inflate(in, out)
	assert.NoError(t, err)
	for _, e := range strings.Split(inflated, "\n") {
		assert.Contains(t, out.String(), e)
	}
}
