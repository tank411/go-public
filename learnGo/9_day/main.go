/*
 * @Descripttion: 模块信息介绍
 * @Company: HIK
 * @version: 1.0.0
 * @Author: TK
 * @Date: 2019-09-25 14:15:58
 * @LastEditors: TK
 * @LastEditTime: 2019-10-10 16:42:44
 */
package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/buguang01/util"
)

//数据映射接口
type IStructSql struct {
	Timeout    int //超时时间（秒）
	InitNum    int //初始化内存空间
	ChanNum    int //通道缓存空间
	SubChanNum int //子通道缓存空间
}

//生成更新SQL
func MarshalUpSql(v interface{}, tablename string) (sql string) {
	result := util.NewStringBuilder()
	result.Append("INSERT INTO ")
	result.Append(tablename)
	result.Append("(")
	t := reflect.TypeOf(v)
	farr := t.Elem()
	tmp := util.NewStringBuilder()
	vtmp := util.NewStringBuilder()
Fieldfor:
	for i := 0; i < farr.NumField(); i++ {
		field := farr.Field(i)
		bigetag := field.Tag.Get("bige")
		narr := strings.Split(bigetag, ",")
		name := field.Name
		iskey := false
		for _, v := range narr {
			switch v {
			case "bigekey":
				iskey = true
			case "select":
			case "-":
				continue Fieldfor
			default:
				name = v
			}
		}
		if !tmp.IsEmpty() {
			result.Append(",")
			tmp.Append(",")
		}
		result.Append(name)
		tmp.Append("?")
		if !iskey {
			if !vtmp.IsEmpty() {
				vtmp.Append(",")
			}
			vtmp.Append(name)
			vtmp.Append("=values(")
			vtmp.Append(name)
			vtmp.Append(")")
		}
	}
	result.Append(")VALUES(")
	result.Append(tmp.ToString())
	result.Append(") ON DUPLICATE KEY UPDATE ")
	result.Append(vtmp.ToString())
	result.Append(";")
	return result.ToString()
}

//生成查询SQL
func MarshalQsql(v interface{}, tablename string) (sql string) {
	result := util.NewStringBuilder()
	result.Append("SELECT ")

	t := reflect.TypeOf(v)
	farr := t.Elem()
	where := util.NewStringBuilder()
Fieldfor:
	for i := 0; i < farr.NumField(); i++ {
		if i > 0 {
			result.Append(",")
		}
		field := farr.Field(i)
		bigetag := field.Tag.Get("bige")
		narr := strings.Split(bigetag, ",")
		name := field.Name
		iswhere := false
		for _, v := range narr {
			switch v {
			case "bigekey":
			case "select":
				iswhere = true
			case "-":
				continue Fieldfor
			default:
				name = v
			}
		}
		result.Append(name)
		if iswhere {
			if !where.IsEmpty() {
				where.Append(" AND ")
			}
			where.Append(name)
			where.Append("=?")
		}
	}
	result.Append(" FROM ")
	result.Append(tablename)
	if !where.IsEmpty() {
		result.Append(" WHERE ")
		result.Append(where.ToString())
	}
	result.Append(";")
	return result.ToString()
}

func main() {
	i := new(IStructSql)
	sqlstr := MarshalQsql(i, "ceshimingcheng")
	fmt.Println(sqlstr)
}
