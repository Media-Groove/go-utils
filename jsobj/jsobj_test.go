package jsobj

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewObj(t *testing.T) {
	json := map[string]any{
		"aaa": float64(1),
		"bbb": "CCC",
	}
	obj := newObj(json)
	assert.Equal(t, "$", obj.Path)
	assert.Equal(t, json, obj.Value)
	assert.Nil(t, obj.Error)
}

func TestJsobj_readMap_正常系(t *testing.T) {
	json := map[string]any{
		"aaa": float64(1),
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	obj := newObj(json)
	obj.readMap("aaa")
	assert.Equal(t, "$.aaa", obj.Path)
	assert.Equal(t, float64(1), obj.Value)
	assert.Nil(t, obj.Error)
}

func TestJsobj_readMap_存在しないキー(t *testing.T) {
	json := map[string]any{
		"aaa": float64(1),
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	obj := newObj(json)
	obj.readMap("AAA")
	assert.Equal(t, "$.AAA", obj.Path)
	assert.Nil(t, obj.Value)
	assert.Nil(t, obj.Error)
}

func TestJsobj_readMap_Mapではない(t *testing.T) {
	json := "string"
	obj := newObj(json)
	obj.readMap("aaa")
	assert.Equal(t, "$", obj.Path)
	assert.Equal(t, json, obj.Value)
	assert.Equal(t, errors.New("not map: $"), obj.Error)
}

func TestJsobj_readArray_正常系(t *testing.T) {
	json := []any{1, true, map[string]any{
		"aaa": "bbb",
	}}
	obj := newObj(json)
	obj.readArray(1)
	assert.Equal(t, "$[1]", obj.Path)
	assert.Equal(t, true, obj.Value)
	assert.Nil(t, obj.Error)
}

func TestJsobj_readArray_存在しないindex(t *testing.T) {
	json := []any{1, true, map[string]any{
		"aaa": "bbb",
	}}
	obj := newObj(json)
	obj.readArray(5)
	assert.Equal(t, "$[5]", obj.Path)
	assert.Nil(t, obj.Value)
	assert.Nil(t, obj.Error)
}

func TestJsobj_readArray_存在しないindex2(t *testing.T) {
	json := []any{1, true, map[string]any{
		"aaa": "bbb",
	}}
	obj := newObj(json)
	obj.readArray(-1)
	assert.Equal(t, "$[-1]", obj.Path)
	assert.Nil(t, obj.Value)
	assert.Nil(t, obj.Error)
}

func TestJsobj_readArray_Arrayではない(t *testing.T) {
	json := map[string]any{
		"aaa": float64(1),
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	obj := newObj(json)
	obj.readArray(0)
	assert.Equal(t, "$", obj.Path)
	assert.Equal(t, json, obj.Value)
	assert.Equal(t, errors.New("not array: $"), obj.Error)
}

func TestJsonobj_read_正常系(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	obj := newObj(json)
	obj.read([]any{"bbb", "ccc"})
	assert.Equal(t, "$.bbb.ccc", obj.Path)
	assert.Equal(t, true, obj.Value)
	assert.Nil(t, obj.Error)

	obj = newObj(json)
	obj.read([]any{"aaa", 2})
	assert.Equal(t, "$.aaa[2]", obj.Path)
	assert.Equal(t, "AAA", obj.Value)
	assert.Nil(t, obj.Error)
}

func TestJsonobj_read_パスにintとstring以外(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	obj := newObj(json)
	obj.read([]any{"bbb", false, 3})
	assert.Equal(t, "$.bbb", obj.Path)
	assert.Equal(t, json["bbb"], obj.Value)
	assert.Equal(t, errors.New("path error"), obj.Error)
}

func TestJsonobj_read_存在しないパス(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	obj := newObj(json)
	obj.read([]any{"bbb", "ccc", "ddd"})
	assert.Equal(t, "$.bbb.ccc", obj.Path)
	assert.Equal(t, true, obj.Value)
	assert.Equal(t, errors.New("not map: $.bbb.ccc"), obj.Error)
}

func TestReadString_正常系(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	s, err := ReadString(json, "bbb", "ddd")
	require.Nil(t, err)
	assert.Equal(t, "EEE", s)
}

func TestReadString_型ミスマッチ(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	s, err := ReadString(json, "bbb", "ccc")
	assert.Equal(t, "", s)
	assert.Equal(t, errors.New("not string: $.bbb.ccc"), err)
}

func TestReadString_存在しないパス(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	s, err := ReadString(json, "bbb", "ccc", "ddd")
	assert.Equal(t, "", s)
	assert.Equal(t, errors.New("not map: $.bbb.ccc"), err)
}

func TestReadFloat64_正常系(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	f, err := ReadFloat64(json, "aaa", 1)
	require.Nil(t, err)
	assert.Equal(t, float64(1), f)
}

func TestReadFloat64_型ミスマッチ(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	f, err := ReadFloat64(json, "aaa", 2)
	assert.Equal(t, float64(0), f)
	assert.Equal(t, errors.New("not float64: $.aaa[2]"), err)
}

func TestReadFloat64_存在しないパス(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	f, err := ReadFloat64(json, "bbb", "ccc", "ddd")
	assert.Equal(t, float64(0), f)
	assert.Equal(t, errors.New("not map: $.bbb.ccc"), err)
}

func TestReadBool_正常系(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	b, err := ReadBool(json, "aaa", 0)
	require.Nil(t, err)
	assert.Equal(t, true, b)
}

func TestReadBool_型ミスマッチ(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	b, err := ReadBool(json, "aaa", 2)
	assert.Equal(t, false, b)
	assert.Equal(t, errors.New("not bool: $.aaa[2]"), err)
}

func TestReadBool_存在しないパス(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	b, err := ReadBool(json, "bbb", "ccc", "ddd")
	assert.Equal(t, false, b)
	assert.Equal(t, errors.New("not map: $.bbb.ccc"), err)
}

func TestReadMap_正常系(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	m, err := ReadMap(json, "bbb")
	require.Nil(t, err)
	assert.Equal(t, json["bbb"], m)
}

func TestReadMap_型ミスマッチ(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	m, err := ReadMap(json, "aaa")
	assert.Nil(t, m)
	assert.Equal(t, errors.New("not map: $.aaa"), err)
}

func TestReadMap_存在しないパス(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	m, err := ReadMap(json, "bbb", "ccc", "ddd")
	assert.Nil(t, m)
	assert.Equal(t, errors.New("not map: $.bbb.ccc"), err)
}

func TestReadArray_正常系(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	a, err := ReadArray(json, "aaa")
	require.Nil(t, err)
	assert.Equal(t, json["aaa"], a)
}

func TestReadArray_型ミスマッチ(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	a, err := ReadArray(json, "bbb")
	assert.Nil(t, a)
	assert.Equal(t, errors.New("not array: $.bbb"), err)
}

func TestReadArray_存在しないパス(t *testing.T) {
	json := map[string]any{
		"aaa": []any{
			true,
			float64(1),
			"AAA",
		},
		"bbb": map[string]any{
			"ccc": true,
			"ddd": "EEE",
		},
	}
	a, err := ReadArray(json, "bbb", "ccc", "ddd")
	assert.Nil(t, a)
	assert.Equal(t, errors.New("not map: $.bbb.ccc"), err)
}
