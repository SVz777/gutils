/**
 * @file    ext.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2020/10/21
 * @desc
 */
package maphelper

import "github.com/SVz777/gutils/json"

type Ext interface {
	Merge(newExtS interface{}) bool
}

func InitExt(e Ext, extS interface{}) bool {
	var data []byte
	switch t := extS.(type) {
	case string:
		if len(t) == 0 {
			return false
		}
		data = []byte(t)
	case []byte:
		if len(t) == 0 {
			return false
		}
		data = t
	default:
		return false
	}
	err := json.Unmarshal(data, &e)
	return err == nil
}

func GetUpdateExt(e Ext, oldExtS string, newExtS string) string {
	if len(newExtS) == 0 {
		return oldExtS
	}
	InitExt(e, oldExtS)
	if !e.Merge(newExtS) {
		return oldExtS
	}

	retExt, _ := json.Marshal(e)
	return string(retExt)
}
