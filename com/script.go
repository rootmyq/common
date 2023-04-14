package com

import (
	"database/sql"
	"fmt"
	"strings"
)

type Field struct {
	TableSchema string
	TableName   string
	ColumnName  string
	DataType    string
}

// CreateGormStructByTable 根据数据库表床创建结构体
func CreateGormStructByTable(config GormConfig, table string) string {

	sqlStr := fmt.Sprintf("select table_schema,table_name,column_name,data_type from information_schema.columns where table_schema='%s' and table_name='%s'", config.DbSchema, table)
	fmt.Println("sqlStr:" + sqlStr)
	dns := config.DbName + ":" + config.DbPassWord + "@tcp(" + config.DbIp + ":" + config.DbPort + ")/" + config.DbSchema
	fmt.Println("dns:" + dns)
	open, err := sql.Open("mysql", dns)
	if err != nil {
		panic("数据库打开异常：" + err.Error())
	}
	defer open.Close()
	rows, err := open.Query(sqlStr)
	if err != nil {
		panic("查询异常:" + err.Error())
	}
	structHead := "type " + StringTool.UnderLineToHump(table, true) + " struct { \n"

	for rows.Next() {
		var TableSchema, TableName, ColumnName, DataType string
		err := rows.Scan(&TableSchema, &TableName, &ColumnName, &DataType)
		if err != nil {
			panic("查询异常:" + err.Error())
		}
		field := Field{TableSchema, TableName, ColumnName, DataType}

		dataType := ""
		if strings.ToLower(field.DataType) == "bigint" {
			dataType = "int"
		} else if strings.ToLower(field.DataType) == "varchar" {
			dataType = "string"
		} else if strings.ToLower(field.DataType) == "timestamp" {
			dataType = "time.Time"
		} else if strings.ToLower(field.DataType) == "decimal" {
			dataType = "float32"
		} else {
			panic("暂时未开放此类型：" + field.DataType)
		}

		structHead = structHead + StringTool.UnderLineToHump(field.ColumnName, true) + " " + dataType + " `json:\"" + StringTool.UnderLineToHump(field.ColumnName, false) + "\"`\n"
	}
	return structHead + " }"
}
