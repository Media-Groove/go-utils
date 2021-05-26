package sliconv

import "strconv"

type Float32 struct {
	Slice []float32
}

func (s Float32) ToBool() ([]bool, error) {
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

func (s Float32) ToFloat32() ([]float32, error) {
	return s.Slice, nil
}

func (s Float32) ToFloat64() ([]float64, error) {
	rs := make([]float64, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = float64(v)
	}
	return rs, nil
}

func (s Float32) ToInt() ([]int, error) {
	rs := make([]int, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int(v)
	}
	return rs, nil
}

func (s Float32) ToInt32() ([]int32, error) {
	rs := make([]int32, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int32(v)
	}
	return rs, nil
}

func (s Float32) ToInt64() ([]int64, error) {
	rs := make([]int64, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int64(v)
	}
	return rs, nil
}

func (s Float32) ToString() ([]string, error) {
	rs := make([]string, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = strconv.FormatFloat(float64(v), 'f', -1, 32)
	}
	return rs, nil
}
