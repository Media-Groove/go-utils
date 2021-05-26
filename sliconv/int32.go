package sliconv

import "strconv"

type Int32 struct {
	Slice []int32
}

func (s Int32) ToBool() []bool {
	rs := make([]bool, len(s.Slice))
	for i, v := range s.Slice {
		if v == 0 {
			rs[i] = false
		} else {
			rs[i] = true
		}
	}
	return rs
}

func (s Int32) ToFloat32() []float32 {
	rs := make([]float32, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = float32(v)
	}
	return rs
}

func (s Int32) ToFloat64() []float64 {
	rs := make([]float64, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = float64(v)
	}
	return rs
}

func (s Int32) ToInt() []int {
	rs := make([]int, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int(v)
	}
	return rs
}

func (s Int32) ToInt64() []int64 {
	rs := make([]int64, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int64(v)
	}
	return rs
}

func (s Int32) ToString() []string {
	rs := make([]string, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = strconv.FormatInt(int64(v), 10)
	}
	return rs
}
