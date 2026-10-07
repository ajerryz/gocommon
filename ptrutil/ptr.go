package ptrutil

func Int(i int) *int {
	return &i
}

func Int32(i int32) *int32 {
	return &i
}

func Int64(i int64) *int64 {
	return &i
}

func Float32(f float32) *float32 {
	return &f
}

func Float64(f float64) *float64 {
	return &f
}

func Str(s string) *string {
	return &s
}

func Bool(b bool) *bool {
	return &b
}
