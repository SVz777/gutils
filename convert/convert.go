/**
 * @file    convert.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2020-03-02
 * @desc
 */
package convert

import (
	"encoding/json"
	"strconv"
)

func Int(v interface{}) (int, bool) {
	if v == nil {
		return 0, false
	}
	switch t := v.(type) {
	case int:
		return t, true
	case int8:
		return int(t), true
	case int16:
		return int(t), true
	case int32:
		return int(t), true
	case int64:
		return int(t), true
	case uint:
		return int(t), true
	case uint8:
		return int(t), true
	case uint16:
		return int(t), true
	case uint32:
		return int(t), true
	case uint64:
		return int(t), true
	case float32:
		return int(t), true
	case float64:
		return int(t), true
	case string:
		if len(t) == 0 {
			return 0, true
		}
		i, err := strconv.Atoi(t)
		return int(i), err == nil
	case json.Number:
		i, err := t.Int64()
		return int(i), err == nil
	case bool:
		if t {
			return 1, true
		}
		return 0, true
	}

	return 0, false
}

func UInt(v interface{}) (uint, bool) {
	if v == nil {
		return 0, false
	}
	switch t := v.(type) {
	case int:
		return uint(t), true
	case int8:
		return uint(t), true
	case int16:
		return uint(t), true
	case int32:
		return uint(t), true
	case int64:
		return uint(t), true
	case uint:
		return t, true
	case uint8:
		return uint(t), true
	case uint16:
		return uint(t), true
	case uint32:
		return uint(t), true
	case uint64:
		return uint(t), true
	case float32:
		return uint(t), true
	case float64:
		return uint(t), true
	case string:
		if len(t) == 0 {
			return 0, true
		}
		i, err := strconv.Atoi(t)
		return uint(i), err == nil
	case json.Number:
		i, err := t.Int64()
		return uint(i), err == nil
	case bool:
		if t {
			return 1, true
		}
		return 0, true
	}

	return 0, false
}

func Int64(v interface{}) (int64, bool) {
	if v == nil {
		return 0, false
	}
	switch t := v.(type) {
	case int:
		return int64(t), true
	case int8:
		return int64(t), true
	case int16:
		return int64(t), true
	case int32:
		return int64(t), true
	case int64:
		return t, true
	case uint:
		return int64(t), true
	case uint8:
		return int64(t), true
	case uint16:
		return int64(t), true
	case uint32:
		return int64(t), true
	case uint64:
		return int64(t), true
	case float32:
		return int64(t), true
	case float64:
		return int64(t), true
	case string:
		if len(t) == 0 {
			return 0, true
		}
		i, err := strconv.ParseInt(t, 10, 64)
		return i, err == nil
	case json.Number:
		i, err := t.Int64()
		return i, err == nil
	case bool:
		if t {
			return 1, true
		}
		return 0, true
	}
	return 0, false
}

func Int32(v interface{}) (int32, bool) {
	switch t := v.(type) {
	case int:
		return int32(t), true
	case int8:
		return int32(t), true
	case int16:
		return int32(t), true
	case int32:
		return t, true
	case int64:
		return int32(t), true
	case uint:
		return int32(t), true
	case uint8:
		return int32(t), true
	case uint16:
		return int32(t), true
	case uint32:
		return int32(t), true
	case uint64:
		return int32(t), true
	case float32:
		return int32(t), true
	case float64:
		return int32(t), true
	case string:
		if len(t) == 0 {
			return 0, true
		}
		i, err := strconv.ParseInt(t, 10, 64)
		return int32(i), err == nil
	case json.Number:
		i, err := t.Int64()
		return int32(i), err == nil
	case bool:
		if t {
			return 1, true
		}
		return 0, true
	}
	return 0, false
}

func Uint64(v interface{}) (uint64, bool) {
	switch t := v.(type) {
	case int:
		return uint64(t), true
	case int8:
		return uint64(t), true
	case int16:
		return uint64(t), true
	case int32:
		return uint64(t), true
	case int64:
		return uint64(t), true
	case uint:
		return uint64(t), true
	case uint8:
		return uint64(t), true
	case uint16:
		return uint64(t), true
	case uint32:
		return uint64(t), true
	case uint64:
		return t, true
	case float32:
		return uint64(t), true
	case float64:
		return uint64(t), true
	case string:
		if len(t) == 0 {
			return 0, true
		}
		i, err := strconv.ParseUint(t, 10, 64)
		return i, err == nil
	case json.Number:
		i, err := t.Int64()
		return uint64(i), err == nil
	case bool:
		if t {
			return 1, true
		}
		return 0, true
	}
	return 0, false
}

func String(v interface{}) (string, bool) {
	switch t := v.(type) {
	case int:
		return strconv.Itoa(t), true
	case int8:
		return strconv.FormatInt(int64(t), 10), true
	case int16:
		return strconv.FormatInt(int64(t), 10), true
	case int32:
		return strconv.FormatInt(int64(t), 10), true
	case int64:
		return strconv.FormatInt(t, 10), true
	case uint:
		return strconv.FormatUint(uint64(t), 10), true
	case uint8:
		return strconv.FormatUint(uint64(t), 10), true
	case uint16:
		return strconv.FormatUint(uint64(t), 10), true
	case uint32:
		return strconv.FormatUint(uint64(t), 10), true
	case uint64:
		return strconv.FormatUint(t, 10), true
	case float32:
		return strconv.FormatFloat(float64(t), 'E', -1, 32), true
	case float64:
		return strconv.FormatFloat(t, 'E', -1, 64), true
	case string:
		return t, true
	case json.Number:
		return t.String(), true
	case bool:
		if t {
			return "true", true
		}
		return "false", true
	case map[string]interface{}:
		if b, err := json.Marshal(t); err != nil {
			return "", false
		} else {
			return string(b), true
		}
	}

	return "", false
}

func Float64(v interface{}) (float64, bool) {
	switch t := v.(type) {
	case int:
		return float64(t), true
	case int8:
		return float64(t), true
	case int16:
		return float64(t), true
	case int32:
		return float64(t), true
	case int64:
		return float64(t), true
	case uint:
		return float64(t), true
	case uint8:
		return float64(t), true
	case uint16:
		return float64(t), true
	case uint32:
		return float64(t), true
	case uint64:
		return float64(t), true
	case float32:
		return float64(t), true
	case float64:
		return t, true
	case string:
		if len(t) == 0 {
			return 0, true
		}
		f, err := strconv.ParseFloat(t, 64)
		return f, err == nil
	case json.Number:
		f, err := t.Float64()
		return f, err == nil
	case bool:
		if t {
			return 1, true
		}
		return 0, true
	}

	return 0, false
}
