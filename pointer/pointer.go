/**
 * @file    helper.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2020-02-06
 * @desc
 */
package pointer

func IntPtr(i int) *int {
	return &i
}

func Int8Ptr(i int8) *int8 {
	return &i
}

func Int16Ptr(i int16) *int16 {
	return &i
}

func Int32Ptr(i int32) *int32 {
	return &i
}

func Int64Ptr(u int64) *int64 {
	return &u
}

func StringPtr(s string) *string {
	return &s
}

func Float64Ptr(f float64) *float64 {
	return &f
}

func Uint64Ptr(u uint64) *uint64 {
	return &u
}
func Uint8Ptr(u uint8) *uint8 {
	return &u
}

func BoolPtr(u bool) *bool {
	return &u
}
