package sliconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool_ToFloat32(t *testing.T) {
	fs := Bool{[]bool{true, false}}.ToFloat32()
	assert.Equal(t, []float32{1, 0}, fs)
}

func TestBool_ToFloat64(t *testing.T) {
	fs := Bool{[]bool{true, false}}.ToFloat64()
	assert.Equal(t, []float64{1, 0}, fs)
}

func TestBool_ToInt(t *testing.T) {
	ns := Bool{[]bool{true, false}}.ToInt()
	assert.Equal(t, []int{1, 0}, ns)
}

func TestBool_ToInt32(t *testing.T) {
	ns := Bool{[]bool{true, false}}.ToInt32()
	assert.Equal(t, []int32{1, 0}, ns)
}

func TestBool_ToInt64(t *testing.T) {
	ns := Bool{[]bool{true, false}}.ToInt64()
	assert.Equal(t, []int64{1, 0}, ns)
}

func TestBool_ToString(t *testing.T) {
	ss := Bool{[]bool{true, false}}.ToString()
	assert.Equal(t, []string{"true", "false"}, ss)
}
