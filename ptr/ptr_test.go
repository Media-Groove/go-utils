package ptr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStr(t *testing.T) {
	p := Str("AAA")
	assert.Equal(t, "AAA", *p)
}

func TestBool(t *testing.T) {
	p := Bool(true)
	assert.Equal(t, true, *p)
}

func TestInt(t *testing.T) {
	p := Int(123)
	assert.Equal(t, 123, *p)
}

func TestInt32(t *testing.T) {
	p := Int32(123)
	assert.Equal(t, int32(123), *p)
}

func TestInt64(t *testing.T) {
	p := Int64(123)
	assert.Equal(t, int64(123), *p)
}

func TestFloat32(t *testing.T) {
	p := Float32(1.23)
	assert.Equal(t, float32(1.23), *p)
}

func TestFloat64(t *testing.T) {
	p := Float64(1.23)
	assert.Equal(t, float64(1.23), *p)
}
