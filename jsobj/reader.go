package jsobj

import (
	"encoding/json"
	"errors"
	"io"
)

type Reader interface {
	ReadString(path ...any) string
	ReadFloat64(path ...any) float64
	ReadInt64(path ...any) int64
	ReadInt(path ...any) int
	ReadBool(path ...any) bool
	ReadMap(path ...any) map[string]any
	ReadArray(path ...any) []any
	Error() error
}

type reader struct {
	obj any
	err error
}

func NewReader(r io.Reader) (Reader, error) {
	bs, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	// オブジェクト
	obj := map[string]any{}
	err = json.Unmarshal(bs, &obj)
	if err == nil {
		return &reader{
			obj: obj,
		}, nil
	}
	// 配列
	a := []any{}
	err = json.Unmarshal(bs, &a)
	if err == nil {
		return &reader{
			obj: a,
		}, nil
	}
	// 文字列
	str := ""
	err = json.Unmarshal(bs, &str)
	if err == nil {
		return &reader{
			obj: str,
		}, nil
	}
	// 数値
	num := float64(0)
	err = json.Unmarshal(bs, &num)
	if err == nil {
		return &reader{
			obj: num,
		}, nil
	}
	return nil, err
}

func (r *reader) ReadString(path ...any) string {
	v, err := ReadString(r.obj, path...)
	if err != nil {
		r.err = errors.Join(r.err, err)
	}
	return v
}

func (r *reader) ReadFloat64(path ...any) float64 {
	v, err := ReadFloat64(r.obj, path...)
	if err != nil {
		r.err = errors.Join(r.err, err)
	}
	return v
}

func (r *reader) ReadInt64(path ...any) int64 {
	v, err := ReadFloat64(r.obj, path...)
	if err != nil {
		r.err = errors.Join(r.err, err)
	}
	return int64(v)
}

func (r *reader) ReadInt(path ...any) int {
	v, err := ReadFloat64(r.obj, path...)
	if err != nil {
		r.err = errors.Join(r.err, err)
	}
	return int(v)
}

func (r *reader) ReadBool(path ...any) bool {
	v, err := ReadBool(r.obj, path...)
	if err != nil {
		r.err = errors.Join(r.err, err)
	}
	return v
}

func (r *reader) ReadMap(path ...any) map[string]any {
	v, err := ReadMap(r.obj, path...)
	if err != nil {
		r.err = errors.Join(r.err, err)
	}
	return v
}

func (r *reader) ReadArray(path ...any) []any {
	v, err := ReadArray(r.obj, path...)
	if err != nil {
		r.err = errors.Join(r.err, err)
	}
	return v
}

func (r *reader) Error() error {
	return r.err
}
