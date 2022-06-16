package array

import (
	"reflect"

	"github.com/SVz777/gutils/collections/set"
)

func ArrayUnion(arrs ...[]int64) []int64 {
	if len(arrs) == 0 {
		return nil
	} else if len(arrs) == 1 {
		return arrs[0]
	}
	values := set.NewInt64Set()
	for _, arr := range arrs {
		for _, v := range arr {
			values.Add(v)
		}
	}
	return values.AllItems()
}

func ArrayInter(arrs ...[]int64) []int64 {
	if len(arrs) == 0 {
		return nil
	} else if len(arrs) == 1 {
		return arrs[0]
	}
	values := set.NewInt64Set(arrs[0]...)
	for _, arr := range arrs[1:] {
		tSet := set.NewInt64Set(arr...)
		for v := range values {
			if !tSet.IsContain(v) {
				values.Delete(v)
			}
		}
	}
	return values.AllItems()
}

func ArrayInArray(v interface{}, arr interface{}) bool {
	vValue, arrValue := reflect.ValueOf(v), reflect.ValueOf(arr)
	if arrValue.Kind() != reflect.Slice && arrValue.Kind() != reflect.Array {
		return false
	}

	if arrValue.Len() <= 0 {
		return false
	}

	if vValue.Kind() == reflect.Ptr {
		vValue = vValue.Elem()
	}
	arrPtrFlag := false
	if arrValue.Type().Elem().Kind() == reflect.Ptr {
		arrPtrFlag = true
		if vValue.Type() != arrValue.Type().Elem().Elem() {
			return false
		}
	} else {
		if vValue.Type() != arrValue.Type().Elem() {
			return false
		}
	}
	for ii := 0; ii < arrValue.Len(); ii++ {
		if arrPtrFlag {
			if vValue.Interface() == arrValue.Index(ii).Elem().Interface() {
				return true
			}
		} else {
			if vValue.Interface() == arrValue.Index(ii).Interface() {
				return true
			}
		}
	}
	return false
}
