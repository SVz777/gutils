/**
 * @file    json.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2019-06-03
 * @desc
 */
package json

import (
	"bytes"
	bjson "encoding/json"
)

type Number = bjson.Number

func Marshal(v interface{}) ([]byte, error) {
	return bjson.Marshal(v)
}

// 处理json float64精度丢失
func Unmarshal(data []byte, v interface{}) error {
	decoder := bjson.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	err := decoder.Decode(&v)
	return err
}
