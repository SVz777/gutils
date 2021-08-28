/**
 * @file    goid.go
 * @author  903943711@qq.com
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * @date    2021/8/28
 * @desc
 */
package goid

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

const idLen = 10 + 20 // +支持goid长度

func GetGoID() int

func SlowGetGoID() int {
	var buf [idLen]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}

	return id

}
