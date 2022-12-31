package ghcomp

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInflater_Inflate(t *testing.T) {
	in := bytes.NewBufferString(deflated)
	out := bytes.NewBuffer(make([]byte, 0))

	err := Inflate(in, out)
	assert.NoError(t, err)

	inflatedElements := strings.Split(inflated, "\n")
	assert.Equal(t, len(strings.Split(out.String(), "\n")), len(inflatedElements))
	for _, e := range inflatedElements {
		assert.Contains(t, out.String(), e)
	}
}
