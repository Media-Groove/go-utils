package jsobj

import (
	"errors"
	"strconv"
)

func ReadString(json any, path ...any) (string, error) {
	obj := newObj(json)
	obj.read(path)
	if obj.Error != nil {
		return "", obj.Error // 存在しない場合
	}
	if s, ok := obj.Value.(string); ok {
		return s, nil // 正常
	}
	// 型ミスマッチ
	return "", errors.New("not string: " + obj.Path)
}

func ReadFloat64(json any, path ...any) (float64, error) {
	obj := newObj(json)
	obj.read(path)
	if obj.Error != nil {
		return 0, obj.Error // 存在しない場合
	}
	if f, ok := obj.Value.(float64); ok {
		return f, nil // 正常
	}
	// 型ミスマッチ
	return 0, errors.New("not float64: " + obj.Path)
}

func ReadBool(json any, path ...any) (bool, error) {
	obj := newObj(json)
	obj.read(path)
	if obj.Error != nil {
		return false, obj.Error // 存在しない場合
	}
	if b, ok := obj.Value.(bool); ok {
		return b, nil // 正常
	}
	// 型ミスマッチ
	return false, errors.New("not bool: " + obj.Path)
}

func ReadMap(json any, path ...any) (map[string]any, error) {
	obj := newObj(json)
	obj.read(path)
	if obj.Error != nil {
		return nil, obj.Error // 存在しない場合
	}
	if m, ok := obj.Value.(map[string]any); ok {
		return m, nil // 正常
	}
	// 型ミスマッチ
	return nil, errors.New("not map: " + obj.Path)
}

func ReadArray(json any, path ...any) ([]any, error) {
	obj := newObj(json)
	obj.read(path)
	if obj.Error != nil {
		return nil, obj.Error // 存在しない場合
	}
	if a, ok := obj.Value.([]any); ok {
		return a, nil // 正常
	}
	// 型ミスマッチ
	return nil, errors.New("not array: " + obj.Path)
}

type jsobj struct {
	Path  string
	Value any
	Error error
}

func newObj(obj any) *jsobj {
	return &jsobj{
		Path:  "$",
		Value: obj,
	}
}

func (obj *jsobj) read(path []any) {
	for _, p := range path {
		if obj.Error != nil {
			break
		}
		switch v := p.(type) {
		case int:
			obj.readArray(v)
		case string:
			obj.readMap(v)
		default:
			obj.Error = errors.New("path error")
		}
	}
}

func (obj *jsobj) readMap(key string) {
	if m, ok := obj.Value.(map[string]any); ok {
		obj.Path += "." + key
		if v, ok := m[key]; ok {
			obj.Value = v
		} else {
			obj.Value = nil
			obj.Error = errors.New("not found: " + obj.Path)
		}
	} else {
		obj.Error = errors.New("not map: " + obj.Path)
	}
}

func (obj *jsobj) readArray(i int) {
	if a, ok := obj.Value.([]any); ok {
		obj.Path += "[" + strconv.Itoa(i) + "]"
		if i >= 0 && i < len(a) {
			obj.Value = a[i]
		} else {
			obj.Value = nil
			obj.Error = errors.New("not found: " + obj.Path)
		}
	} else {
		obj.Error = errors.New("not array: " + obj.Path)
	}
}
