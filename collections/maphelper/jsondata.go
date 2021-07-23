/**
 * @file    jsondata.go
 * @author
 *  ___  _  _  ____
 * / __)( \/ )(_   )
 * \__ \ \  /  / /_
 * (___/  \/  (____)
 * (903943711@qq.com)
 * @date    2020-05-29
 * @desc
 */
package maphelper

type JsonData = map[string]interface{}

func JDGet(jd JsonData, ks []string) (ret interface{}, ok bool) {
	if len(ks) == 1 {
		ret, ok = jd[ks[0]]
	} else {
		if nx, ok1 := jd[ks[0]]; ok1 {
			if nxx, ok1 := nx.(JsonData); ok1 {
				ret, ok = JDGet(nxx, ks[1:])
			} else {
				ret, ok = nx, true
			}
		} else {
			ret, ok = nil, false
		}

	}
	return
}

func JDSet(jd JsonData, ks []string, value interface{}) bool {
	if len(ks) == 1 {
		jd[ks[0]] = value
	} else {
		if nx, ok1 := jd[ks[0]]; ok1 {
			if nxx, ok1 := nx.(JsonData); ok1 {
				JDSet(nxx, ks[1:], value)
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

func JDSetWithFunc(jd JsonData, ks []string, f func(interface{}) (interface{}, bool)) bool {
	if len(ks) == 1 {
		if v, ok := f(jd[ks[0]]); ok {
			jd[ks[0]] = v
		} else {
			return false
		}
	} else {
		if nx, ok1 := jd[ks[0]]; ok1 {
			if nxx, ok1 := nx.(JsonData); ok1 {
				JDSetWithFunc(nxx, ks[1:], f)
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

// 获取嵌套map，非嵌套map误用
func JDGetOne(jd JsonData) JsonData {
	for _, v := range jd {
		if vv, ok := v.(JsonData); ok {
			return vv
		}
	}
	return nil
}
