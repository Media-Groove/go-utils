package sliconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt32_ToBool(t *testing.T) {
	bs := Int32{[]int32{1, 2, 0, -1}}.ToBool()
	assert.Equal(t, []bool{true, true, false, true}, bs)
}

func TestInt32_ToFloat32(t *testing.T) {
	fs := Int32{[]int32{1, 2, 0, -1}}.ToFloat32()
	assert.Equal(t, []float32{1, 2, 0, -1}, fs)
}

func TestInt32_ToFloat64(t *testing.T) {
	fs := Int32{[]int32{1, 2, 0, -1}}.ToFloat64()
	assert.Equal(t, []float64{1, 2, 0, -1}, fs)
}

func TestInt32_ToInt(t *testing.T) {
	ns := Int32{[]int32{1, 2, 0, -1}}.ToInt()
	assert.Equal(t, []int{1, 2, 0, -1}, ns)
}

func TestInt32_ToInt64(t *testing.T) {
	ns := Int32{[]int32{1, 2, 0, -1}}.ToInt64()
	assert.Equal(t, []int64{1, 2, 0, -1}, ns)
}

func TestInt32_ToString(t *testing.T) {
	ns := Int32{[]int32{1, 2, 0, -1}}.ToString()
	assert.Equal(t, []string{"1", "2", "0", "-1"}, ns)
}
