package sliconv

import "strconv"

type Int struct {
	Slice []int
}

func (s Int) ToBool() []bool {
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

func (s Int) ToFloat32() []float32 {
	rs := make([]float32, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = float32(v)
	}
	return rs
}

func (s Int) ToFloat64() []float64 {
	rs := make([]float64, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = float64(v)
	}
	return rs
}

func (s Int) ToInt32() []int32 {
	rs := make([]int32, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int32(v)
	}
	return rs
}

func (s Int) ToInt64() []int64 {
	rs := make([]int64, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = int64(v)
	}
	return rs
}

func (s Int) ToString() []string {
	rs := make([]string, len(s.Slice))
	for i, v := range s.Slice {
		rs[i] = strconv.Itoa(v)
	}
	return rs
}
