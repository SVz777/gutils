/**
 * @file    jsonpath.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2021-08-10
 * @desc
 */
package json

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/SVz777/gutils/convert"
)

const jsonPathTag = "json_path"

// JSONPath json的tag封装
type JSONPath struct {
	data interface{}
}

// NewJSONPath ...
func NewJSONPath(jsonData []byte) (*JSONPath, error) {
	j := new(JSONPath)
	err := j.UnmarshalJSON(jsonData)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// MarshalJSON json.Marshaler
func (j *JSONPath) MarshalJSON() ([]byte, error) {
	return Marshal(&j.data)
}

// UnmarshalJSON json.Unmarshaler
func (j *JSONPath) UnmarshalJSON(p []byte) error {
	return Unmarshal(p, &j.data)
}

// IsNil 值是否为空
func (j *JSONPath) IsNil() bool {
	return j.data == nil
}

// Get 获取key 对应值，不存在 IsNil 为 true
func (j *JSONPath) Get(key interface{}) *JSONPath {
	switch data := j.data.(type) {
	case map[string]interface{}:
		k, ok := key.(string)
		if !ok {
			return &JSONPath{data: nil}
		}
		v, ok := data[k]
		if !ok {
			return &JSONPath{data: nil}
		}
		return &JSONPath{data: v}
	case []interface{}:
		k, err := convert.Int(key)
		if err != nil {
			return &JSONPath{data: nil}
		}
		if len(data) <= k {
			return &JSONPath{data: nil}
		}
		return &JSONPath{data: data[k]}
	default:
		return &JSONPath{data: nil}
	}
}

// Get2 获取key 对应值，取到了第二个值为true
func (j *JSONPath) Get2(key interface{}) (*JSONPath, bool) {
	switch data := j.data.(type) {
	case map[string]interface{}:
		k, ok := key.(string)
		if !ok {
			return &JSONPath{data: nil}, false
		}
		v, ok := data[k]
		if !ok {
			return &JSONPath{data: nil}, false
		}
		return &JSONPath{data: v}, true
	case []interface{}:
		k, ok := key.(int)
		if !ok {
			return &JSONPath{data: nil}, false
		}
		if len(data) <= k {
			return &JSONPath{data: nil}, false
		}
		return &JSONPath{data: data[k]}, true
	default:
		return &JSONPath{data: nil}, false
	}
}

// GetPath 根据path 获取
func (j *JSONPath) GetPath(path ...string) *JSONPath {
	t := j
	for _, p := range path {
		t = t.Get(p)
	}
	return t
}

// Interface 获取data值
func (j *JSONPath) Interface() interface{} {
	return j.data
}

// Int 将值转为 int
func (j *JSONPath) Int() (int, error) {
	return convert.Int(j.data)
}

// Int32 将值转为 int32
func (j *JSONPath) Int32() (int32, error) {
	return convert.Int32(j.data)
}

// Int64 将值转为 int64
func (j *JSONPath) Int64() (int64, error) {
	return convert.Int64(j.data)
}

// Uint 将值转为 uint
func (j *JSONPath) Uint() (uint, error) {
	return convert.Uint(j.data)
}

// UInt64 将值转为 Uint64
func (j *JSONPath) UInt64() (uint64, error) {
	return convert.Uint64(j.data)
}

// String 将值转为 string
func (j *JSONPath) String() (string, error) {
	return convert.String(j.data)
}

// Float64 将值转为 float64
func (j *JSONPath) Float64() (float64, error) {
	return convert.Float64(j.data)
}

// Bool 将值转为 bool
func (j *JSONPath) Bool() (bool, error) {
	return convert.Bool(j.data)
}

// Map 将值断言为 map[string]interface
func (j *JSONPath) Map() (map[string]interface{}, error) {
	if m, ok := (j.data).(map[string]interface{}); ok {
		return m, nil
	}
	return nil, fmt.Errorf("type assertion to map[string]interface{} failed")
}

// Array 将值断言为 []interface
func (j *JSONPath) Array() ([]interface{}, error) {
	if a, ok := (j.data).([]interface{}); ok {
		return a, nil
	}
	return nil, fmt.Errorf("type assertion to []interface{} failed")
}

// StringArray 将值转为 []string
func (j *JSONPath) StringArray() ([]string, error) {
	a, err := j.Array()
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0, len(a))
	for idx, s := range a {
		v, err := convert.String(s)
		if err != nil {
			return nil, fmt.Errorf("%v:%v convert err: %w", idx, s, err)
		}
		ret = append(ret, v)
	}
	return ret, nil
}

// ParseWithJSONPath 根据data的tag(json_path)定义来填充data
func (j *JSONPath) ParseWithJSONPath(data interface{}) error {
	tr := reflect.TypeOf(data)
	if tr.Kind() != reflect.Ptr {
		// data必须是指针
		return fmt.Errorf("data must be *struct")
	}
	tr = tr.Elem()
	if tr.Kind() != reflect.Struct {
		// *data必须是结构体
		return fmt.Errorf("data must be *struct")
	}
	vr := reflect.ValueOf(data).Elem()

	fn := tr.NumField()
	for ii := 0; ii < fn; ii++ {
		tf := tr.Field(ii)
		vf := vr.Field(ii)
		busSrc := tf.Tag.Get(jsonPathTag)
		if busSrc == "" {
			if tf.Type.Kind() == reflect.Struct {
				err := j.ParseWithJSONPath(vf.Addr().Interface())
				if err != nil {
					return fmt.Errorf("sub struct %s parse err: %w", tf.Name, err)
				}
			}
			continue
		}
		busPath := strings.Split(busSrc, ".")
		jsonValue := j.GetPath(busPath...)

		value, err := j.getValue(tf.Name, tf.Type, jsonValue)
		if err != nil {
			return fmt.Errorf("getvalue error: %w", err)
		}
		vr.Field(ii).Set(value)
	}
	return nil
}

func (j *JSONPath) getValue(fieldName string, tf reflect.Type, jsonValue *JSONPath) (reflect.Value, error) {
	switch tf.Kind() {
	case reflect.Map:
		v, err := jsonValue.Map()
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(v), nil
	case reflect.Struct:
		trueValue := reflect.New(tf)
		b, _ := jsonValue.MarshalJSON()
		err := json.Unmarshal(b, trueValue.Interface())
		if err != nil {
			return reflect.Value{}, fmt.Errorf("%s json unmarshal err: %w", fieldName, err)
		}
		return trueValue.Elem(), nil
	case reflect.Slice:
		itemKind := tf.Elem().Kind()
		if itemKind == reflect.Interface {
			aValue, err := jsonValue.Array()
			if err != nil {
				return reflect.Value{}, fmt.Errorf("%s :%w", fieldName, err)
			}
			return reflect.ValueOf(aValue), nil
		} else if itemKind == reflect.String {
			aValue, err := jsonValue.StringArray()
			if err != nil {
				return reflect.Value{}, fmt.Errorf("%s :%w", fieldName, err)
			}
			return reflect.ValueOf(aValue), nil
		} else {
			aValue, err := jsonValue.Array()
			if err != nil {
				return reflect.Value{}, fmt.Errorf("%s :%w", fieldName, err)
			}
			trueValue := reflect.MakeSlice(tf, len(aValue), len(aValue))
			for idx := range aValue {
				iv, err := j.getValue(fieldName, tf.Elem(), jsonValue.Get(idx))
				if err != nil {
					return reflect.Value{}, fmt.Errorf("%s parse slice err: %w", fieldName, err)
				}
				trueValue.Index(idx).Set(iv)
			}
			return trueValue, nil
		}
	default:
		iv, err1 := convert.Convert(jsonValue.Interface(), tf.Kind())
		if err1 != nil {
			return reflect.Value{}, fmt.Errorf("%s :%w", fieldName, err1)
		}
		return reflect.ValueOf(iv), nil
	}
}
