package sliconv

type Bool struct {
	Slice []bool
}

func (s Bool) ToFloat32() []float32 {
	rs := make([]float32, len(s.Slice))
	for i, v := range s.Slice {
		if v {
			rs[i] = 1
		} else {
			rs[i] = 0
		}
	}
	return rs
}

func (s Bool) ToFloat64() []float64 {
	rs := make([]float64, len(s.Slice))
	for i, v := range s.Slice {
		if v {
			rs[i] = 1
		} else {
			rs[i] = 0
		}
	}
	return rs
}

func (s Bool) ToInt() []int {
	rs := make([]int, len(s.Slice))
	for i, v := range s.Slice {
		if v {
			rs[i] = 1
		} else {
			rs[i] = 0
		}
	}
	return rs
}

func (s Bool) ToInt32() []int32 {
	rs := make([]int32, len(s.Slice))
	for i, v := range s.Slice {
		if v {
			rs[i] = 1
		} else {
			rs[i] = 0
		}
	}
	return rs
}

func (s Bool) ToInt64() []int64 {
	rs := make([]int64, len(s.Slice))
	for i, v := range s.Slice {
		if v {
			rs[i] = 1
		} else {
			rs[i] = 0
		}
	}
	return rs
}

func (s Bool) ToString() []string {
	rs := make([]string, len(s.Slice))
	for i, v := range s.Slice {
		if v {
			rs[i] = "true"
		} else {
			rs[i] = "false"
		}
	}
	return rs
}
