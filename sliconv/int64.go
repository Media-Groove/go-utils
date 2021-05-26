package sliconv

import "strconv"

type Int64 struct {
	Slice []int64
}

func (s Int64) ToBool() []bool {
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

func (s Int64) ToFloat32() []float32 {
	rs := make([]float32, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = float32(v)
	}
	return rs
}

func (s Int64) ToFloat64() []float64 {
	rs := make([]float64, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = float64(v)
	}
	return rs
}

func (s Int64) ToInt() []int {
	rs := make([]int, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int(v)
	}
	return rs
}

func (s Int64) ToInt32() []int32 {
	rs := make([]int32, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int32(v)
	}
	return rs
}

func (s Int64) ToString() []string {
	rs := make([]string, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = strconv.FormatInt(int64(v), 10)
	}
	return rs
}
