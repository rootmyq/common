package com

import "encoding/json"

type JSONUtil struct {
}

// ToString  对象转json
func (*JSONUtil) ToString(obj any) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic("序列化异常:" + err.Error())
	}
	return string(bytes)
}

// ToObject 字符串传json
func (*JSONUtil) ToObject(str string, obj any) any {
	bytes := []byte(str)
	err := json.Unmarshal(bytes, obj)
	if err != nil {
		panic("json转对象异常")
	}
	return obj
}
