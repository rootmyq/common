package com

import (
	"errors"
	"time"
)

type DateUtil struct {
}

// GetCurrentTime 获取当前时间
func (d *DateUtil) GetCurrentTime(format string) (string, error) {
	var realFormat string
	switch format {
	case "yyyy-MM-dd":
		realFormat = "2006-01-02"
	case "yyyy/MM/dd":
		realFormat = "2006/01/02"
	case "yyyy-MM-dd HH:mm:ss":
		realFormat = "2006-01-02 15:04:05"
	default:
		return "", errors.New("无效时间类型：" + format)
	}
	now := time.Now()
	t := time.Unix(now.Unix(), 0) // 参数分别是：秒数,纳秒数
	return t.Format(realFormat), nil
}
