package com

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

type StringUtils struct {
}

// IsEmpty 判断为空
func (s *StringUtils) IsEmpty(str string) bool {
	return len(str) == 0
}

// IsNotEmpty 判断不为空
func (s *StringUtils) IsNotEmpty(str string) bool {
	return !s.IsEmpty(str)
}

// IsBlank 判断为空(去掉空格之后)
func (s *StringUtils) IsBlank(str string) bool {
	return s.IsEmpty(str) || s.IsEmpty(strings.Trim(str, " "))
}

// IsNotBlank 判断不为空(去掉空格之后)
func (s *StringUtils) IsNotBlank(str string) bool {
	return s.IsNotEmpty(str) && s.IsNotEmpty(strings.Trim(str, " "))
}

// Eq 判断相等
func (s *StringUtils) Eq(str1 string, str2 string) bool {
	return str1 == str2
}

// NotEq 判断不相等
func (s *StringUtils) NotEq(str1 string, str2 string) bool {
	return !s.Eq(str1, str2)
}

// Len 获取长度
func (s *StringUtils) Len(str string) int {
	return len(str)
}

// Split 获取长度
func (s *StringUtils) Split(str string, sep string) []string {
	return strings.Split(str, sep)
}

// Join 拼接字符串
func (s *StringUtils) Join(strs []string, sep string) string {
	return strings.Join(strs, sep)
}

// Contain 是否包含字符串
func (s *StringUtils) Contain(str string, substr string) bool {
	return strings.Contains(str, substr)
}

// ReplaceAll 是否包含字符串
func (s *StringUtils) ReplaceAll(str string, old string, new string) string {
	return strings.ReplaceAll(str, old, new)
}

// HumpToUnderLine 下划线转驼峰
func (s *StringUtils) HumpToUnderLine(str string) string {
	if len(str) == 0 {
		return str
	}
	letterStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var arr []string
	for i := range str {
		ss := str[i]
		if strings.Contains(letterStr, string(ss)) && i != 0 && i != len(str)-1 {
			arr = append(arr, "_")
			arr = append(arr, string(ss))
		} else {
			arr = append(arr, string(ss))
		}
	}

	join := strings.Join(arr, "")
	return strings.ToLower(join)
}

// UnderLineToHump 驼峰转下划线
func (s *StringUtils) UnderLineToHump(str string, firstToUpper bool) string {
	if len(str) == 0 {
		return str
	}
	letterStr := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var arr []string
	flag := false
	for i := range str {
		ss := string(str[i])
		if ss == "_" {
			flag = true
			continue
		}
		if flag && strings.Contains(letterStr, ss) {
			arr = append(arr, strings.ToUpper(ss))
		} else {
			arr = append(arr, strings.ToLower(ss))
		}
		flag = false
	}
	if firstToUpper {
		arr[0] = strings.ToUpper(arr[0])
	} else {
		arr[0] = strings.ToLower(arr[0])
	}
	return strings.Join(arr, "")
}

// CreateUUID 创建uuid
func (s *StringUtils) CreateUUID() string {
	v4 := uuid.NewV4()
	return v4.String()
}
