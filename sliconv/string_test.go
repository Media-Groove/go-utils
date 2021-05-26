package sliconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString_ToBool(t *testing.T) {
	bs, err := String{[]string{"true", "True", "false", "FALSE", "0", "1"}}.ToBool()
	require.Nil(t, err)
	assert.Equal(t, []bool{true, true, false, false, false, true}, bs)
}

func TestString_ToBool_error(t *testing.T) {
	_, err := String{[]string{"true", "false", "ERROR"}}.ToBool()
	require.NotNil(t, err)
	ec, ok := err.(*ConvertError)
	assert.True(t, ok)
	assert.Equal(t, 2, ec.Index)
}

func TestString_ToFloat32(t *testing.T) {
	fs, err := String{[]string{"0", "1", "2.3"}}.ToFloat32()
	require.Nil(t, err)
	assert.Equal(t, []float32{0, 1, 2.3}, fs)
}

func TestString_ToFloat32_error(t *testing.T) {
	_, err := String{[]string{"0", "1", "2.3", "ERROR"}}.ToFloat32()
	require.NotNil(t, err)
	ec, ok := err.(*ConvertError)
	assert.True(t, ok)
	assert.Equal(t, 3, ec.Index)
}

func TestString_ToFloat64(t *testing.T) {
	fs, err := String{[]string{"0", "1", "2.3"}}.ToFloat64()
	require.Nil(t, err)
	assert.Equal(t, []float64{0, 1, 2.3}, fs)
}

func TestString_ToFloat64_error(t *testing.T) {
	_, err := String{[]string{"0", "1", "2.3", "ERROR"}}.ToFloat64()
	require.NotNil(t, err)
	ec, ok := err.(*ConvertError)
	assert.True(t, ok)
	assert.Equal(t, 3, ec.Index)
}

func TestString_ToInt(t *testing.T) {
	ns, err := String{[]string{"-1", "0", "100"}}.ToInt()
	require.Nil(t, err)
	assert.Equal(t, []int{-1, 0, 100}, ns)
}

func TestString_ToInt_error(t *testing.T) {
	_, err := String{[]string{"-1", "0", "100", "ERROR"}}.ToInt()
	require.NotNil(t, err)
	ec, ok := err.(*ConvertError)
	assert.True(t, ok)
	assert.Equal(t, 3, ec.Index)
}

func TestString_ToInt32(t *testing.T) {
	ns, err := String{[]string{"-1", "0", "100"}}.ToInt32()
	require.Nil(t, err)
	assert.Equal(t, []int32{-1, 0, 100}, ns)
}

func TestString_ToInt32_error(t *testing.T) {
	_, err := String{[]string{"-1", "0", "100", "ERROR"}}.ToInt32()
	require.NotNil(t, err)
	ec, ok := err.(*ConvertError)
	assert.True(t, ok)
	assert.Equal(t, 3, ec.Index)
}

func TestString_ToInt64(t *testing.T) {
	ns, err := String{[]string{"-1", "0", "100"}}.ToInt64()
	require.Nil(t, err)
	assert.Equal(t, []int64{-1, 0, 100}, ns)
}

func TestString_ToInt64_error(t *testing.T) {
	_, err := String{[]string{"-1", "0", "100", "ERROR"}}.ToInt64()
	require.NotNil(t, err)
	ec, ok := err.(*ConvertError)
	assert.True(t, ok)
	assert.Equal(t, 3, ec.Index)
}
