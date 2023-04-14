package com

import "os"

type FileUtils struct {
}

// ReadBytes 读取文件到内存
func (f *FileUtils) ReadBytes(fPath string) []byte {
	openFile, err := os.Open(fPath)
	if err != nil {
		panic("文件读取失败：" + err.Error())
	}
	defer openFile.Close()

	stat, err := openFile.Stat()
	if err != nil {
		panic("文件状态读取失败：" + err.Error())
	}

	size := stat.Size()
	buff := make([]byte, 0, size)
	_, err = openFile.Read(buff)
	if err != nil {
		panic("文件读取失败：" + err.Error())
	}

	return buff
}
