package com

import (
	"runtime/debug"
)

// LogStack 日志异常栈
type LogStack struct {
	filters []string
}

/*
*
filter 设置过滤字段
*/
func (s *LogStack) filter(f string) *LogStack {
	s.filters = append(s.filters, f)
	return s
}

// Stack 错误栈信息组装
func (s *LogStack) Stack(msg string) []string {
	stack := debug.Stack()
	str := string(stack)
	stackArr := StringTool.Split(str, "\n")

	var rntStacks []string
	rntStacks = append(rntStacks, msg)
	firstLine := false
	for _, stack := range stackArr {
		flag := false
		for i := range s.filters {
			filterStr := s.filters[i]
			if StringTool.Contain(stack, filterStr) {
				flag = true
				break
			}
		}
		if flag {
			if !firstLine {
				firstLine = true
				continue
			}
			stack = StringTool.ReplaceAll(stack, "\t", "")
			rntStacks = append(rntStacks, stack)
		}
	}
	return rntStacks
}
