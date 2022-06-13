package maphelper

import "github.com/SVz777/gutils/json"

func MapFlip(src map[string]string) map[string]string {
	newMap := make(map[string]string, len(src))
	for k, v := range src {
		newMap[v] = k
	}
	return newMap
}

func MapCopy(src map[string]string) map[string]string {
	newMap := make(map[string]string, len(src))
	for k, v := range src {
		newMap[k] = v
	}
	return newMap
}

func StructToMap(obj interface{}) map[string]interface{} {
	data := make(map[string]interface{})
	databyte, _ := json.Marshal(obj)
	_ = json.Unmarshal(databyte, &data)
	return data
}

func MapCopyInterface(src map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{}, len(src))
	for k, v := range src {
		newMap[k] = v
	}
	return newMap
}
