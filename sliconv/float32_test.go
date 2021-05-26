package sliconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFloat32_ToBool(t *testing.T) {
	bs, err := Float32{[]float32{1, 2, 0, -1}}.ToBool()
	require.Nil(t, err)
	assert.Equal(t, []bool{true, true, false, true}, bs)
}

func TestFloat32_ToFloat32(t *testing.T) {
	fs, err := Float32{[]float32{1, 2, 0, -1}}.ToFloat32()
	require.Nil(t, err)
	assert.Equal(t, []float32{1, 2, 0, -1}, fs)
}

func TestFloat32_ToFloat64(t *testing.T) {
	fs, err := Float32{[]float32{1, 2, 0, -1}}.ToFloat64()
	require.Nil(t, err)
	assert.Equal(t, []float64{1, 2, 0, -1}, fs)
}

func TestFloat32_ToInt(t *testing.T) {
	ns, err := Float32{[]float32{1, 2, 0, -1, 1.1, 2.9, -1.9}}.ToInt()
	require.Nil(t, err)
	assert.Equal(t, []int{1, 2, 0, -1, 1, 2, -1}, ns)
}

func TestFloat32_ToInt32(t *testing.T) {
	ns, err := Float32{[]float32{1, 2, 0, -1, 1.1, 2.9, -1.9}}.ToInt32()
	require.Nil(t, err)
	assert.Equal(t, []int32{1, 2, 0, -1, 1, 2, -1}, ns)
}

func TestFloat32_ToInt64(t *testing.T) {
	ns, err := Float32{[]float32{1, 2, 0, -1, 1.1, 2.9, -1.9}}.ToInt64()
	require.Nil(t, err)
	assert.Equal(t, []int64{1, 2, 0, -1, 1, 2, -1}, ns)
}

func TestFloat32_ToString(t *testing.T) {
	ns, err := Float32{[]float32{1, 2, 0, -1, 1.1}}.ToString()
	require.Nil(t, err)
	assert.Equal(t, []string{"1", "2", "0", "-1", "1.1"}, ns)
}
