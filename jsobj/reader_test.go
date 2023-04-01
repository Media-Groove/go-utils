package jsobj

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	// 数値
	r, err := NewReader(strings.NewReader("1"))
	require.Nil(t, err)
	assert.Equal(t, float64(1), r.(*reader).obj)
	// 文字列
	r, err = NewReader(strings.NewReader(`"TEST"`))
	require.Nil(t, err)
	assert.Equal(t, "TEST", r.(*reader).obj)
	// 配列
	r, err = NewReader(strings.NewReader(`[1, false, "AAA"]`))
	require.Nil(t, err)
	assert.Equal(t, []any{float64(1), false, "AAA"}, r.(*reader).obj)
	// オブジェクト
	r, err = NewReader(strings.NewReader(`{"AAA": "BBB"}`))
	require.Nil(t, err)
	assert.Equal(t, map[string]any{"AAA": "BBB"}, r.(*reader).obj)
}

func TestReader_Readいろいろ(t *testing.T) {
	json := `
{
	"AAA": {
		"BBB": "CCC"
	},
	"DDD": [
		1,
		true
	]
}
`
	r, err := NewReader(strings.NewReader(json))
	require.Nil(t, err)
	assert.Equal(t, map[string]any{
		"BBB": "CCC",
	}, r.ReadMap("AAA"))
	assert.Equal(t, []any{float64(1), true}, r.ReadArray("DDD"))
	assert.Equal(t, "CCC", r.ReadString("AAA", "BBB"))
	assert.Equal(t, float64(1), r.ReadFloat64("DDD", 0))
	assert.Equal(t, int64(1), r.ReadInt64("DDD", 0))
	assert.Equal(t, 1, r.ReadInt("DDD", 0))
	assert.Equal(t, true, r.ReadBool("DDD", 1))
	assert.Nil(t, r.Error())
	assert.Nil(t, r.ReadMap("abc"))
	assert.NotNil(t, r.Error())
	assert.Equal(t, "CCC", r.ReadString("AAA", "BBB"))
	assert.NotNil(t, r.Error())
}
