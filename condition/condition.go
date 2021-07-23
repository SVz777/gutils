/**
 * @file    Condition.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2020/8/11
 * @desc
 */
package condition

import (
	"github.com/SVz777/gutils/convert"
	"reflect"
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
