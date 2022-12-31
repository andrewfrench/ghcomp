package ghcomp

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	precision = 12
	inflated  = "bdvkhunfnc90\nbdvkj7vyvtz5\nbdvkjkjbpr1t\nbdvkjkjbremh\nbdvkjkn00jdn\n"
	deflated  = "bdvkhunfnc90\nj7vyvtz5\nkjbpr1t\nremh\nn00jdn\n"
)

func TestTree_Entree(t *testing.T) {
	tree := New(3)

	err := tree.Entree([]byte("1234"))
	assert.ErrorIs(t, err, ErrPrecision)

	err = tree.Entree([]byte("123"))
	assert.NoError(t, err)

	err = tree.Entree([]byte("12"))
	assert.ErrorIs(t, err, ErrPrecision)
}

func TestTree_EntreeDeflated(t *testing.T) {
	in := bytes.NewBufferString(deflated)
	out := bytes.NewBuffer(nil)
	expectedElements := strings.Split(inflated, "\n")

	tree := New(precision)
	err := tree.EntreeDeflated(in)
	assert.NoError(t, err)

	err = tree.WriteInflated(out)
	assert.NoError(t, err)
	for _, e := range expectedElements {
		assert.Contains(t, out.String(), e)
	}

	newElement := "1234567890ab"
	err = tree.Entree([]byte(newElement))
	assert.NoError(t, err)

	out = bytes.NewBuffer(nil)
	err = tree.WriteInflated(out)
	assert.NoError(t, err)
	for _, e := range append(expectedElements, newElement) {
		assert.Contains(t, out.String(), e)
	}
}
