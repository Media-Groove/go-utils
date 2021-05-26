package sliconv

import "fmt"

type ConvertError struct {
	Index int
	Err   error
}

func (e *ConvertError) Error() string {
	return fmt.Sprintf("convert: index %d: %s", e.Index, e.Err.Error())
}

func NewConvertError(i int, err error) *ConvertError {
	return &ConvertError{
		Index: i,
		Err:   err,
	}
}
