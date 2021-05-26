package sliconv

import "strconv"

type Float64 struct {
	Slice []float64
}

func (s Float64) ToBool() ([]bool, error) {
	rs := make([]bool, len(s.Slice))
	for i, v := range s.Slice {
		if v == 0 {
			rs[i] = false
		} else {
			rs[i] = true
		}
	}
	return rs, nil
}

func (s Float64) ToFloat32() ([]float32, error) {
	rs := make([]float32, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = float32(v)
	}
	return rs, nil
}

func (s Float64) ToFloat64() ([]float64, error) {
	return s.Slice, nil
}

func (s Float64) ToInt() ([]int, error) {
	rs := make([]int, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int(v)
	}
	return rs, nil
}

func (s Float64) ToInt32() ([]int32, error) {
	rs := make([]int32, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int32(v)
	}
	return rs, nil
}

func (s Float64) ToInt64() ([]int64, error) {
	rs := make([]int64, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int64(v)
	}
	return rs, nil
}

func (s Float64) ToString() ([]string, error) {
	rs := make([]string, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = strconv.FormatInt(int64(v), 10)
	}
	return rs, nil
}
