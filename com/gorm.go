package com

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"reflect"
	"time"
)

type GormLocalTime time.Time

// MarshalJSON 重写读取时间党发
func (t *GormLocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t GormLocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *GormLocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = GormLocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// GormConfig gorm配置
type GormConfig struct {
	DbIp       string
	DbPort     string
	DbSchema   string
	DbName     string
	DbPassWord string
}

type Gorm struct {
}

// InitGorm 初始化gorm
func (*Gorm) InitGorm(config GormConfig) *gorm.DB {
	//获得一个*grom.DB对象
	DB, err := gorm.Open(
		mysql.New(mysql.Config{
			DSN: "" + config.DbName +
				":" + config.DbPassWord + "@tcp(" + config.DbIp + ":" + config.DbPort + ")/" + config.DbSchema +
				"?charset=utf8&parseTime=True&loc=Local", // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		panic("Gorm 异常：" + err.Error())
	}
	//根据*grom.DB对象获得*sql.DB的通用数据库接口
	sqlDb, err := DB.DB()
	if err != nil {
		panic("Gorm 异常：" + err.Error())
	}
	//defer sqlDb.Close()
	//sqlDb.SetMaxIdleConns(database.MaxConn) //设置最大连接数
	//sqlDb.SetMaxOpenConns(database.MaxOpen) //设置最大的空闲连接数
	data, _ := json.Marshal(sqlDb.Stats()) //获得当前的SQL配置情况
	fmt.Println(string(data))
	return DB
}

// SqlGenerate sql生成器
type SqlGenerate struct {
}

func (g *SqlGenerate) GenerateQuerySql(tableObj any) {

	r := reflect.TypeOf(tableObj)
	//tableName :=StringTool.HumpToUnderLine(r.Name())
	v := reflect.ValueOf(tableObj)
	vMap := make(map[int]reflect.Value)
	for i := 0; i < v.NumField(); i++ {
		vMap[i] = v.Field(i)
	}

	var dbFields []string
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		fmt.Println(field.Type.Kind().String())
		dbFields = append(dbFields, StringTool.HumpToUnderLine(field.Name))
	}
}
