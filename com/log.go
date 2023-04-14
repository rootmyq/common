package com

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strconv"
)

// LogConfig 日志配置
type LogConfig struct {
	LogPath        string
	LogProjectName string
	LogIsDebug     string
}

type Log struct {
	log *zap.Logger
}

// InitLog 初始化日志
func (l *Log) InitLog(config LogConfig) {

	// 此处的配置是从我的项目配置文件读取的，读者可以根据自己的情况来设置
	logPath := config.LogPath
	name := config.LogProjectName
	debug := config.LogIsDebug

	if logPath == "" {
		panic("配置LogPath为空")
	}
	if name == "" {
		panic("配置LogProjectName为空")
	}
	if debug == "" {
		panic("配置LogIsDebug为空")
	}

	hook := lumberjack.Logger{
		Filename:   logPath, // 日志文件路径
		MaxSize:    1,       // 每个日志文件保存的大小 单位:M
		MaxAge:     7,       // 文件最多保存多少天
		MaxBackups: 30,      // 日志文件最多保存多少个备份
		Compress:   false,   // 是否压缩
	}
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // 短路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zap.DebugLevel)
	var writes = []zapcore.WriteSyncer{zapcore.AddSync(&hook)}
	// 如果是开发环境，同时在控制台上也输出
	debugBool, err := strconv.ParseBool(debug)
	if err != nil {
		panic("非bool值【" + debug + "】：" + err.Error())
	}

	if debugBool {
		writes = append(writes, zapcore.AddSync(os.Stdout))
	}
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(writes...),
		atomicLevel,
	)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	field := zap.Fields(zap.String("appName", name))

	// 构造日志
	l.log = zap.New(core, caller, development, field)
	fmt.Println("日志初始化成功")
}

func (l *Log) GetLog() *zap.Logger {
	if l.log == nil {
		panic("日志对象未初始化")
	}
	return l.log
}
