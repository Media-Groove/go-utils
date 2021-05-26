package sliconv

import "strconv"

type String struct {
	Slice []string
}

func (s String) ToBool() ([]bool, error) {
	var err error
	rs := make([]bool, len(s.Slice))
	for i, v := range s.Slice {
		rs[i], err = strconv.ParseBool(v)
		if err != nil {
			return nil, NewConvertError(i, err)
		}
	}
	return rs, nil
}

func (s String) ToFloat32() ([]float32, error) {
	rs := make([]float32, len(s.Slice))
	for i, v := range s.Slice {
		f, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return nil, NewConvertError(i, err)
		}
		rs[i] = float32(f)
	}
	return rs, nil
}

func (s String) ToFloat64() ([]float64, error) {
	rs := make([]float64, len(s.Slice))
	for i, v := range s.Slice {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, NewConvertError(i, err)
		}
		rs[i] = float64(f)
	}
	return rs, nil
}

func (s String) ToInt() ([]int, error) {
	var err error
	rs := make([]int, len(s.Slice))
	for i, v := range s.Slice {
		rs[i], err = strconv.Atoi(v)
		if err != nil {
			return nil, NewConvertError(i, err)
		}
	}
	return rs, nil
}

func (s String) ToInt32() ([]int32, error) {
	rs := make([]int32, len(s.Slice))
	for i, v := range s.Slice {
		n, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, NewConvertError(i, err)
		}
		rs[i] = int32(n)
	}
	return rs, nil
}

func (s String) ToInt64() ([]int64, error) {
	var err error
	rs := make([]int64, len(s.Slice))
	for i, v := range s.Slice {
		rs[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, NewConvertError(i, err)
		}
	}
	return rs, nil
}
