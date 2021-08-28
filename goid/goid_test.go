/**
 * @file    goid_test.go
 * @author  903943711@qq.com
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * @date    2021/8/28
 * @desc
 */
package goid_test

import (
	"testing"

	"github.com/SVz777/gutils/goid"
)

func BenchmarkSlowGetGoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goid.SlowGetGoID()
	}
}

func BenchmarkGetGoID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = goid.GetGoID()
	}
}
