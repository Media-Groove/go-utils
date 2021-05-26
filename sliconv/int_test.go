package sliconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt_ToBool(t *testing.T) {
	bs := Int{[]int{1, 2, 0, -1}}.ToBool()
	assert.Equal(t, []bool{true, true, false, true}, bs)
}

func TestInt_ToFloat32(t *testing.T) {
	fs := Int{[]int{1, 2, 0, -1}}.ToFloat32()
	assert.Equal(t, []float32{1, 2, 0, -1}, fs)
}

func TestInt_ToFloat64(t *testing.T) {
	fs := Int{[]int{1, 2, 0, -1}}.ToFloat64()
	assert.Equal(t, []float64{1, 2, 0, -1}, fs)
}

func TestInt_ToInt32(t *testing.T) {
	ns := Int{[]int{1, 2, 0, -1}}.ToInt32()
	assert.Equal(t, []int32{1, 2, 0, -1}, ns)
}

func TestInt_ToInt64(t *testing.T) {
	ns := Int{[]int{1, 2, 0, -1}}.ToInt64()
	assert.Equal(t, []int64{1, 2, 0, -1}, ns)
}

func TestInt_ToString(t *testing.T) {
	ns := Int{[]int{1, 2, 0, -1}}.ToString()
	assert.Equal(t, []string{"1", "2", "0", "-1"}, ns)
}
