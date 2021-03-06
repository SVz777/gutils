package condition

import (
	"reflect"

	"github.com/SVz777/gutils/convert"
)

func Equal(a interface{}, b interface{}) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na == nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa == sb
	}
	return false
}

func Gt(a interface{}, b interface{}) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na > nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa > sb
	}
	return false
}

func Gte(a interface{}, b interface{}) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na >= nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa >= sb
	}
	return false
}

func Lt(a interface{}, b interface{}) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na < nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa < sb
	}
	return false
}

func Lte(a interface{}, b interface{}) bool {
	na, err1 := convert.Float64(a)
	nb, err2 := convert.Float64(b)
	if err1 != nil && err2 != nil {
		return na <= nb
	}
	sa, err1 := convert.String(a)
	sb, err2 := convert.String(b)
	if err1 != nil && err2 != nil {
		return sa <= sb
	}
	return false
}

func InArr(a interface{}, b interface{}) bool {
	tb := reflect.ValueOf(b)
	if tb.Kind() == reflect.Slice || tb.Kind() == reflect.Array {
		for ii := tb.Len() - 1; ii >= 0; ii-- {
			if Equal(a, tb.Index(ii).Interface()) {
				return true
			}
		}
	}
	return false
}
