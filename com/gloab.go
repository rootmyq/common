package com

// FileTool 获取文件处理工具
var FileTool = new(FileUtils)

// StringTool 获取文件处理工具
var StringTool = new(StringUtils)

// JsonTool 获取文件处理工具
var JsonTool = new(JSONUtil)

// DateTool 时间工具
var DateTool = new(DateUtil)

// LogStackTool 只获取当前项目的异常栈
var LogStackTool = new(LogStack).filter("go_base")
