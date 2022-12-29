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
	for _, e := range strings.Split(inflated, "\n") {
		assert.Contains(t, out.String(), e)
	}
}
