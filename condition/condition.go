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
	na, ok1 := convert.Float64(a)
	nb, ok2 := convert.Float64(b)
	if ok1 && ok2 {
		return na == nb
	}
	sa, ok1 := convert.String(a)
	sb, ok2 := convert.String(b)
	if ok1 && ok2 {
		return sa == sb
	}
	return false
}

func Gt(a interface{}, b interface{}) bool {
	na, ok1 := convert.Float64(a)
	nb, ok2 := convert.Float64(b)
	if ok1 && ok2 {
		return na > nb
	}
	sa, ok1 := convert.String(a)
	sb, ok2 := convert.String(b)
	if ok1 && ok2 {
		return sa > sb
	}
	return false
}

func Gte(a interface{}, b interface{}) bool {
	na, ok1 := convert.Float64(a)
	nb, ok2 := convert.Float64(b)
	if ok1 && ok2 {
		return na >= nb
	}
	sa, ok1 := convert.String(a)
	sb, ok2 := convert.String(b)
	if ok1 && ok2 {
		return sa >= sb
	}
	return false
}

func Lt(a interface{}, b interface{}) bool {
	na, ok1 := convert.Float64(a)
	nb, ok2 := convert.Float64(b)
	if ok1 && ok2 {
		return na < nb
	}
	sa, ok1 := convert.String(a)
	sb, ok2 := convert.String(b)
	if ok1 && ok2 {
		return sa < sb
	}
	return false
}

func Lte(a interface{}, b interface{}) bool {
	na, ok1 := convert.Float64(a)
	nb, ok2 := convert.Float64(b)
	if ok1 && ok2 {
		return na <= nb
	}
	sa, ok1 := convert.String(a)
	sb, ok2 := convert.String(b)
	if ok1 && ok2 {
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
