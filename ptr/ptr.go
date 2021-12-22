package ptr

// Str 指定した値のポインタを取得する
func Str(s string) *string {
	return &s
}

// Bool 指定した値のポインタを取得する
func Bool(b bool) *bool {
	return &b
}

// Int 指定した値のポインタを取得する
func Int(n int) *int {
	return &n
}

// Int32 指定した値のポインタを取得する
func Int32(n int32) *int32 {
	return &n
}

// Int64 指定した値のポインタを取得する
func Int64(n int64) *int64 {
	return &n
}

// Float32 指定した値のポインタを取得する
func Float32(n float32) *float32 {
	return &n
}

// Float64 指定した値のポインタを取得する
func Float64(n float64) *float64 {
	return &n
}
